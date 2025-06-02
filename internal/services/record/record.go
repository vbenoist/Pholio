package record

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackidu14/pholio/internal/database/connector"
	apimodels "github.com/jackidu14/pholio/internal/models/api"
	databasemodels "github.com/jackidu14/pholio/internal/models/database"
	mongopagination "github.com/jackidu14/pholio/pkg/helpers/database"
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

func GetRecentlyRecords(pgParams apimodels.PaginationQuery) (*apimodels.PaginatedResult[apimodels.RecentlyRecords], error) {
	var results apimodels.PaginatedResult[apimodels.RecentlyRecords]

	var records []databasemodels.Record
	preparedQuery := getBasePreparedQuery(pgParams)
	paginatedData, err := (*preparedQuery).Filter(bson.D{}).Decode(&records).Find()
	if err != nil {
		return &results, err
	}
	if len(records) == 0 {
		return &results, nil
	}

	var lastRecord []databasemodels.Record
	lastRecordPrpQ := getBasePreparedQuery(apimodels.PaginationQuery{SortBy: "date", SortAsc: -1, PerPage: 1, Page: 1})
	_, err = (*lastRecordPrpQ).Filter(bson.D{}).Decode(&lastRecord).Find()
	if err != nil {
		return &results, err
	}

	refinedRecords, err := extractRecentlyRecords(records, lastRecord[0])
	if err != nil {
		return &results, err
	}

	results.Pagination = paginatedData.Pagination
	results.Document = *refinedRecords
	return &results, nil
}

func extractRecentlyRecords(records []databasemodels.Record, referenceRecord databasemodels.Record) (*apimodels.RecentlyRecords, error) {
	/* 2 weeks before last records */
	dateDelimiter := referenceRecord.Date.Time().AddDate(0, 0, -14) // 24 * 14

	fmt.Printf("Date to delimit: %s\n", dateDelimiter)
	var recentlyRecords apimodels.RecentlyRecords

	for _, rec := range records {
		if rec.Date.Time().Before(dateDelimiter) {
			fmt.Printf("Adding %s in lately\n", rec.Date.Time())
			recentlyRecords.Lately = append(recentlyRecords.Lately, rec)
		} else {
			fmt.Printf("Adding %s in lastly\n", rec.Date.Time())
			recentlyRecords.Lastly = append(recentlyRecords.Lastly, rec)
		}
	}

	return &recentlyRecords, nil
}

func GetRecordsGroupByDate(pgParams apimodels.PaginationQuery) (*apimodels.PaginatedResults[apimodels.GroupbyRecord], error) {
	var results apimodels.PaginatedResults[apimodels.GroupbyRecord]

	if pgParams.SortBy == "" || pgParams.SortBy == "date" {
		pgParams.SortBy = "_id"
	}
	preparedQuery := getBasePreparedQuery(pgParams)
	aggPaginatedData, err := (*preparedQuery).Aggregate(
		bson.M{
			"$sort": bson.M{pgParams.SortBy: pgParams.SortAsc},
		},
		bson.M{
			"$limit": pgParams.PerPage,
		},
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

func GetRecordsGroupByDateDebug(pgParams apimodels.PaginationQuery) (*apimodels.PaginatedResults[apimodels.GroupbyRecord], error) {
	var results apimodels.PaginatedResults[apimodels.GroupbyRecord]

	collection := connector.GetCollection("records")
	if pgParams.SortBy == "" || pgParams.SortBy == "date" {
		pgParams.SortBy = "_id"
	}

	// cursor, err := collection.Aggregate(context.Background(), []bson.M{
	// 	{
	// 		"$sort": bson.M{"_id": 1},
	// 	},
	// 	{
	// 		"$limit": pgParams.PerPage,
	// 	},
	// 	{
	// 		"$group": bson.M{
	// 			"_id": bson.M{"date": bson.M{"$dateTrunc": bson.M{
	// 				"date": "$date",
	// 				"unit": "day",
	// 			}}},
	// 			"records": bson.M{"$push": "$$ROOT"},
	// 			"total":   bson.M{"$sum": 1},
	// 		},
	// 	},
	// })

	cursor, err := collection.Aggregate(context.Background(), []bson.M{
		{
			"$facet": bson.M{
				"data": []bson.M{
					{"$sort": bson.M{"_id": 1}},
					{"$limit": pgParams.PerPage},
				},
				"total": []bson.M{{"$count": "count"}},
			},
		},
		{
			"$group": bson.M{
				"_id": bson.M{"date": bson.M{"$dateTrunc": bson.M{
					"date": "$date",
					"unit": "day",
				}}},
				"records": bson.M{"$push": "$$ROOT"},
				"total":   bson.M{"$sum": 1},
			},
		},
		// {
		// 	bson.M{"$facet": bson.M{
		// 		"data":  facetData,
		// 		"total": []bson.M{{"$count": "count"}},
		// 	},
		// },
	})

	if err != nil {
		return &results, err
	}
	defer cursor.Close(context.Background())

	var ress []bson.M
	if err = cursor.All(context.TODO(), &ress); err != nil {
		log.Fatal(err)
	}

	if len(ress) == 0 {
		return &results, nil
	}

	fmt.Printf("\n\n######\nALL RESULTS: (%d) \n", len(ress))
	for _, result := range ress {
		fmt.Println(result)
	}
	fmt.Print("\n######DONE######\n\n")

	return &results, nil
}

func GetRecordsGroupByLocation(pgParams apimodels.PaginationQuery) (*apimodels.PaginatedResults[apimodels.GroupbyRecord], error) {
	var results apimodels.PaginatedResults[apimodels.GroupbyRecord]

	if pgParams.SortBy == "" || pgParams.SortBy == "date" {
		pgParams.SortBy = "_id"
	}
	preparedQuery := getBasePreparedQuery(pgParams)

	aggPaginatedData, err := (*preparedQuery).Aggregate(
		bson.M{
			"$sort": bson.M{pgParams.SortBy: pgParams.SortAsc},
		},
		bson.M{
			"$limit": pgParams.PerPage,
		},
		bson.M{
			"$group": bson.M{"_id": "$location", "records": bson.M{"$push": "$$ROOT"}},
		},
	)
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
