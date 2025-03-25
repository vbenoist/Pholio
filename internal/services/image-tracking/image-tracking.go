package imagetracking

import (
	"context"
	"strings"

	"github.com/jackidu14/pholio/internal/database/connector"
	"github.com/jackidu14/pholio/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetRecordImageTracking(recordId string) (models.RecordImageTracking, error) {
	var result models.RecordImageTracking
	collection := connector.GetCollection("image-status")
	rcid, _ := primitive.ObjectIDFromHex(recordId)

	filter := bson.D{primitive.E{Key: "record-id", Value: rcid}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	return result, err
}

func AddRecordImageTracking(recordId string) (*mongo.InsertOneResult, error) {
	collection := connector.GetCollection("image-status")
	rcid, _ := primitive.ObjectIDFromHex(recordId)

	imageTracking := models.RecordImageTracking{RecordId: rcid}
	return collection.InsertOne(context.TODO(), imageTracking)
}

func UpdateRecordImageTracking(recordId string, fields bson.D) (*mongo.UpdateResult, error) {
	collection := connector.GetCollection("image-status")
	rcid, _ := primitive.ObjectIDFromHex(recordId)
	formattedFields := fields

	for idx := range fields {
		formattedFields[idx].Key = strings.ToLower(fields[idx].Key)
	}

	filter := bson.D{primitive.E{Key: "recordid", Value: rcid}}
	update := bson.D{primitive.E{Key: "$set", Value: formattedFields}}

	return collection.UpdateOne(context.TODO(), filter, update)
}
