package record

import (
	"context"

	mongopagination "github.com/gobeam/mongo-go-pagination"
	"github.com/jackidu14/pholio/internal/database/connector"
	apimodels "github.com/jackidu14/pholio/internal/models/api"
	databasemodels "github.com/jackidu14/pholio/internal/models/database"
	"go.mongodb.org/mongo-driver/bson"
)

func GetRecords() (apimodels.PaginatedResults[databasemodels.Record], error) {
	var documents []databasemodels.Record
	var results apimodels.PaginatedResults[databasemodels.Record]
	collection := connector.GetCollection("records")

	// paginatedData, err := mongopagination.New(collection).Context(context.Background()).Limit(10).Page(1).Find()
	paginatedData, err := mongopagination.New(collection).Context(context.Background()).Filter(bson.D{}).Limit(10).Page(1).Decode(&documents).Find()
	if err != nil {
		return results, err
	}

	results.Pagination = paginatedData.Pagination
	results.Documents = documents
	return results, nil
}
