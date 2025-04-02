package record

import (
	"context"

	"github.com/jackidu14/pholio/internal/database/connector"
	databasemodels "github.com/jackidu14/pholio/internal/models/database"
	"go.mongodb.org/mongo-driver/bson"
)

func GetRecords() ([]databasemodels.Record, error) {
	// /* TODO : Add filter for max date */
	// /* TODO : Add pagination */
	var documents []databasemodels.Record
	collection := connector.GetCollection("records")

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return documents, err
	}
	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &documents)
	return documents, err
}
