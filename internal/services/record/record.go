package record

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/vbenoist/pholio/internal/database/connector"
	apimodels "github.com/vbenoist/pholio/internal/models/api"
	databasemodels "github.com/vbenoist/pholio/internal/models/database"
	mongopagination "github.com/vbenoist/pholio/pkg/helpers/database"
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

func GetRecordsStars(pgParams apimodels.PaginationQuery) (*apimodels.PaginatedResults[databasemodels.DetailedRecord], error) {
	var results apimodels.PaginatedResults[databasemodels.DetailedRecord]
	preparedQuery := getBasePreparedQuery(pgParams)

	aggPaginatedData, err := (*preparedQuery).Aggregate(
		bson.M{
			"$lookup": bson.M{
				"from":         "pin",
				"as":           "record_pin",
				"localField":   "_id",
				"foreignField": "record_id",
				"pipeline": []bson.M{
					{
						"$project": bson.M{
							"pin_id": "$_id",
						},
					},
					{
						"$project": bson.M{
							"_id":       0,
							"record_id": 0,
						},
					},
				},
			},
		},
	)
	if err != nil {
		return &results, err
	}

	if len(aggPaginatedData.Data) == 0 {
		return &results, nil
	}

	results.Pagination = aggPaginatedData.Pagination
	extracted := struct {
		Record  *databasemodels.Record
		Details *databasemodels.DetailedRecord
	}{}

	for _, bsonRaw := range aggPaginatedData.Data {
		if marshallErr := bson.Unmarshal(bsonRaw, &extracted.Record); marshallErr != nil {
			continue
		}
		if marshallErr := bson.Unmarshal(bsonRaw, &extracted.Details); marshallErr != nil {
			continue
		}

		results.Documents = append(results.Documents, databasemodels.DetailedRecord{
			Record: *(extracted.Record),
			Pin:    *&(extracted.Details).Pin,
		})
	}

	// collection := connector.GetCollection("records")
	// cursor, err := collection.Aggregate(context.Background(), []bson.M{
	// 	{
	// 		"$lookup": bson.M{
	// 			"from":         "pin",
	// 			"as":           "record_pin",
	// 			"localField":   "_id",
	// 			"foreignField": "record_id",
	// 			"pipeline": []bson.M{
	// 				// {
	// 				// 	"$replaceRoot": bson.M{
	// 				// 		"newRoot": bson.M{
	// 				// 			"pin": bson.M{"$cond": bson.M{
	// 				// 				"if":   bson.M{"$eq": []interface{}{"", "$record_id"}},
	// 				// 				"then": "0",
	// 				// 				"else": "1",
	// 				// 			}},
	// 				// 		},
	// 				// 	},
	// 				// },
	// 				{
	// 					"$project": bson.M{
	// 						"pin_id": "$_id",
	// 					},
	// 				},
	// 				{
	// 					"$project": bson.M{
	// 						"_id":       0,
	// 						"record_id": 0,
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// 	// {
	// 	// 	"$replaceRoot": bson.M{
	// 	// 		"newRoot": bson.M{
	// 	// 			"$mergeObjects": []interface{}{
	// 	// 				// bson.M{
	// 	// 				// 	"$arrayElemAt": []interface{}{"$isPin", 0},
	// 	// 				// },
	// 	// 				bson.M{
	// 	// 					// "pin": bson.M{
	// 	// 					// 	"$cond": bson.M{
	// 	// 					// 		"if":   bson.M{"$eq": []interface{}{"1", "$isPin.pin"}},
	// 	// 					// 		"then": "1",
	// 	// 					// 		"else": "0",
	// 	// 					// 	},
	// 	// 					// },
	// 	// 					"pin": bson.M{
	// 	// 						"$ifNull": []string{"$isPin.pin", "0"},
	// 	// 					},
	// 	// 				},
	// 	// 				"$$ROOT",
	// 	// 			},
	// 	// 		},
	// 	// 	},
	// 	// },
	// 	// {
	// 	// 	"$project": bson.M{
	// 	// 		// "pin": "$record_id",
	// 	// 		"isPin": 0,
	// 	// 		// "$$ROOT": 1,
	// 	// 		// "pin": bson.M{"$cond": []bson.M{
	// 	// 		// 	{"if": bson.M{"$eq": []interface{}{"", "$record_id"}}},
	// 	// 		// 	{"then": "0"},
	// 	// 		// 	{"else": "1"},
	// 	// 		// }},
	// 	// 		// "pin": bson.M{"$ifNull": []bson.M{
	// 	// 		// 	{"if": bson.M{"$eq": []interface{}{"", "$record_id"}}},
	// 	// 		// 	{"then": "0"},
	// 	// 		// 	{"else": "1"},
	// 	// 		// }},
	// 	// 		// "isPin": 0,
	// 	// 	},
	// 	// },
	// })

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

	if pgParams.SortBy == "" {
		pgParams.SortBy = "date"
	}
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

func GetRecordsGroupByDateDebug(pgParams apimodels.PaginationQuery) (*apimodels.PaginatedResults[apimodels.GroupbyRecord], error) {
	var results apimodels.PaginatedResults[apimodels.GroupbyRecord]

	collection := connector.GetCollection("records")
	if pgParams.SortBy == "" {
		pgParams.SortBy = "date"
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

	fmt.Printf("Skip: %d\n", pgParams.PerPage*(pgParams.Page-1))
	cursor, err := collection.Aggregate(context.Background(), []bson.M{
		// {"$sort": bson.M{"_id": 1}},
		// {"$skip": pgParams.PerPage * (pgParams.Page - 1)},
		// {"$limit": pgParams.PerPage},
		// {
		// 	"$group": bson.M{
		// 		"_id": bson.M{"date": bson.M{"$dateTrunc": bson.M{
		// 			"date": "$date",
		// 			"unit": "day",
		// 		}}},
		// 		"lines": bson.M{"$push": "$$ROOT"},
		// 		"total": bson.M{"$sum": 1},
		// 	},
		// },

		// bson.D{
		// {
		// 	Key: "$facet", Value: bson.M{
		// 		"data": bson.D{
		// 			/* Sorting, skip & limit on ungroupped lines, before applying groupBy */
		// 			{Key: "$sort", Value: bson.M{"date": pgParams.SortAsc}},
		// 			{Key: "$skip", Value: pgParams.PerPage * (pgParams.Page - 1)},
		// 			{Key: "$limit", Value: pgParams.PerPage},
		// 			/* Applying groupBy */
		// 			{
		// 				Key: "$group", Value: bson.M{
		// 					"_id": bson.M{"date": bson.M{"$dateTrunc": bson.M{
		// 						"date": "$date",
		// 						"unit": "day",
		// 					}}},
		// 					"lines": bson.M{"$push": "$$ROOT"},
		// 					"total": bson.M{"$sum": 1},
		// 				},
		// 			},
		// 			/* Sorting the groupped output */
		// 			{Key: "$sort", Value: bson.M{"_id": pgParams.SortAsc}},
		// 		},
		// 		"total": []bson.M{{"$count": "count"}}, // In the facet, outside limit scope to count every db lines
		// 	},
		// },

		{
			"$facet": bson.M{
				"data": []bson.M{
					/* Sorting, skip & limit on ungroupped lines, before applying groupBy */
					{"$sort": bson.M{"date": pgParams.SortAsc}},
					{"$skip": pgParams.PerPage * (pgParams.Page - 1)},
					{"$limit": pgParams.PerPage},
					/* Applying groupBy */
					{
						"$group": bson.M{
							"_id": bson.M{"date": bson.M{"$dateTrunc": bson.M{
								"date": "$date",
								"unit": "day",
							}}},
							"lines": bson.M{"$push": "$$ROOT"},
							"total": bson.M{"$sum": 1},
						},
					},
					/* Sorting the groupped output */
					{"$sort": bson.M{"_id": pgParams.SortAsc}},
				},
				"total": []bson.M{{"$count": "count"}}, // In the facet, outside limit scope to count every db lines
			},
		},

		// {"$skip": pgParams.PerPage * (pgParams.Page - 1)},
		// {"$limit": pgParams.PerPage},
		// {"$sort": bson.M{"date": pgParams.SortAsc}},
		// {
		// 	"$group": bson.M{
		// 		"_id": bson.M{"date": bson.M{"$dateTrunc": bson.M{
		// 			"date": "$date",
		// 			"unit": "day",
		// 		}}},
		// 		"lines": bson.M{"$push": "$$ROOT"},
		// 		"total": bson.M{"$sum": 1},
		// 	},
		// },
		// { // capturing groupped data through facet and working on it
		// 	"$facet": bson.M{
		// 		"data": []bson.M{
		// 			{"$sort": bson.M{"date": pgParams.SortAsc}},
		// 		},
		// 		"total": []bson.M{{"$count": "count"}}, // In the facet, outside limit scope to count every db lines
		// 	},
		// },
	})

	// test := bson.M{
	// 	"$group": bson.M{
	// 		"_id": bson.M{"date": bson.M{"$dateTrunc": bson.M{
	// 			"date": "$date",
	// 			"unit": "day",
	// 		}}},
	// 		"lines": bson.M{"$push": "$$ROOT"},
	// 		"total": bson.M{"$sum": 1},
	// 	},
	// }

	// data, err := bson.Marshal(test)
	// if err != nil {
	// 	panic(err)
	// }

	// decoder, err := bson.NewDecoder(bsonrw.NewBSONDocumentReader(data))
	// if err != nil {
	// 	panic(err)
	// }

	// var res interface{}
	// err = decoder.Decode(&res)
	// if err != nil {
	// 	panic(err)
	// }

	// str := fmt.Sprintf("%+v\n", res)
	// fmt.Printf("%s\n", str)

	// // isGroup := strings.Contains(str, "$group")

	// // rest, err := bson.MarshalExtJSON(test, true, true)
	// // if err != nil {
	// // 	return &results, err
	// // }

	// fmt.Printf("Test: %s\n", test)

	// cursor, err := collection.Aggregate(context.Background(), []bson.M{
	// 	{"$sort": bson.M{"_id": 1}},
	// 	{"$skip": pgParams.PerPage * (pgParams.Page - 1)},
	// 	{"$limit": pgParams.PerPage},
	// 	{
	// 		"$group": bson.M{
	// 			"_id": bson.M{"date": bson.M{"$dateTrunc": bson.M{
	// 				"date": "$date",
	// 				"unit": "day",
	// 			}}},
	// 			"lines": bson.M{"$push": "$$ROOT"},
	// 			"total": bson.M{"$sum": 1},
	// 		},
	// 	},
	// 	// {
	// 	// 	"$facet": bson.M{
	// 	// 		"data": []bson.M{
	// 	// 			{"$sort": bson.M{"_id": 1}},
	// 	// 			{"$skip": pgParams.PerPage * (pgParams.Page - 1)},
	// 	// 			{"$limit": pgParams.PerPage},
	// 	// 		},
	// 	// 		"total": []bson.M{{"$count": "count"}},
	// 	// 	},
	// 	// },
	// 	// {
	// 	// 	bson.M{"$facet": bson.M{
	// 	// 		"data":  facetData,
	// 	// 		"total": []bson.M{{"$count": "count"}},
	// 	// 	},
	// 	// },
	// })

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
	if pgParams.SortBy == "" {
		pgParams.SortBy = "location"
	}
	preparedQuery := getBasePreparedQuery(pgParams)

	/* Adding a second sort, to get in group-by results ordered by date */
	(*preparedQuery).Sort("date", -1)

	aggPaginatedData, err := (*preparedQuery).Aggregate(
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

	if pgParams.SortAscGroup != 0 {
		preparedQuery = preparedQuery.SortGroup(pgParams.SortAscGroup)
	} else {
		preparedQuery = preparedQuery.SortGroup(-1)
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
