package record

import (
	"context"
	"time"

	mongopagination "github.com/gobeam/mongo-go-pagination"
	"github.com/jackidu14/pholio/internal/database/connector"
	apimodels "github.com/jackidu14/pholio/internal/models/api"
	databasemodels "github.com/jackidu14/pholio/internal/models/database"
	"go.mongodb.org/mongo-driver/bson"
)

func GetRecords(pgParams apimodels.PaginationQuery) (*apimodels.PaginatedResults[databasemodels.Record], error) {
	var documents []databasemodels.Record
	var results apimodels.PaginatedResults[databasemodels.Record]

	preparedQuery := getBasePreparedQuery(pgParams)
	paginatedData, err := (*preparedQuery).Filter(bson.D{}).Decode(&documents).Find()
	if err != nil {
		return &results, err
	}

	results.Pagination = paginatedData.Pagination
	results.Documents = documents
	return &results, nil
}

func GetRecordsGroupByDate(pgParams apimodels.PaginationQuery) (*apimodels.PaginatedResults[apimodels.GroupbyRecord], error) {
	var results apimodels.PaginatedResults[apimodels.GroupbyRecord]

	preparedQuery := getBasePreparedQuery(pgParams)
	aggPaginatedData, err := (*preparedQuery).Aggregate(
		bson.M{
			"$group": bson.M{
				"_id": bson.M{"date": bson.M{"$dateTrunc": bson.M{
					"date": "$date",
					"unit": "day",
				}}},
				"records": bson.M{"$push": "$$ROOT"},
			},
		},
	)
	if err != nil {
		return &results, err
	}

	if len(aggPaginatedData.Data) == 0 {
		return &results, nil
	}

	extracted, err := extractGrouppedQuery(aggPaginatedData.Data, func(raw bson.RawValue) (string, error) {
		var ext time.Time

		err = raw.Unmarshal(&ext)
		if err != nil {
			return "", err
		}

		return ext.Format("2006-02-01"), nil
	})
	if err != nil {
		return &results, err
	}

	results.Documents = *extracted
	return &results, nil
}

func GetRecordsGroupByLocation(pgParams apimodels.PaginationQuery) (*apimodels.PaginatedResults[apimodels.GroupbyRecord], error) {
	var results apimodels.PaginatedResults[apimodels.GroupbyRecord]
	preparedQuery := getBasePreparedQuery(pgParams)

	aggPaginatedData, err := (*preparedQuery).Aggregate(bson.M{"$group": bson.M{"_id": "$location", "records": bson.M{"$push": "$$ROOT"}}})
	if err != nil {
		return &results, err
	}

	if len(aggPaginatedData.Data) == 0 {
		return &results, nil
	}

	extracted, err := extractGrouppedQuery(aggPaginatedData.Data, nil)
	if err != nil {
		return &results, err
	}

	results.Documents = *extracted
	return &results, nil
}

func getBasePreparedQuery(pgParams apimodels.PaginationQuery) *mongopagination.PagingQuery {
	collection := connector.GetCollection("records")

	preparedQuery := mongopagination.New(collection).Context(context.Background())
	if pgParams.SortBy != "" {
		preparedQuery = preparedQuery.Sort(pgParams.SortBy, pgParams.SortAsc)
	}

	if pgParams.Page != 0 {
		preparedQuery = preparedQuery.Page(pgParams.Page)
	}

	if pgParams.PerPage != 0 {
		preparedQuery = preparedQuery.Limit(pgParams.PerPage)
	}

	return &preparedQuery
}

func extractGrouppedQuery(rawResults []bson.Raw, customGroupbyUnmarshall func(bson.RawValue) (string, error)) (*[]apimodels.GroupbyRecord, error) {
	var results []apimodels.GroupbyRecord
	var groupRecord apimodels.GroupbyRecord
	var record *databasemodels.Record

	/* For each group by */
	for _, bsonRaw := range rawResults {
		/*
			row is an array with 2 elements:
				- key 0 : group by value
				- key 1 : bson array of grouped values
		*/
		row, err := bsonRaw.Values()
		if err != nil {
			return nil, err
		}

		/* Group by value can be an embeded document, like for date f.ex. */
		if row[0].Type == bson.TypeEmbeddedDocument && customGroupbyUnmarshall != nil {
			embbed, err := row[0].Document().Values()
			if err != nil {
				return nil, err
			}

			res, err := customGroupbyUnmarshall(embbed[0])
			if err != nil {
				return nil, err
			}

			groupRecord = apimodels.GroupbyRecord{
				GroupBy: res,
			}
		} else {
			groupRecord = apimodels.GroupbyRecord{
				GroupBy: row[0].StringValue(),
			}
		}

		/* Trying to get array of grouped values from bson */
		arrEls, err := row[1].Array().Values()
		if err != nil {
			return nil, err
		}

		for _, el := range arrEls {
			if marshallErr := bson.Unmarshal(el.Value, &record); marshallErr == nil {
				groupRecord.Results = append(groupRecord.Results, *record)
			} else {
				return nil, err
			}
		}

		results = append(results, groupRecord)
	}

	return &results, nil
}
