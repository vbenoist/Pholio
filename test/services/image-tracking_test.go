package services_test

import (
	"context"
	"testing"

	"github.com/jackidu14/pholio/internal/database/connector"
	"github.com/jackidu14/pholio/internal/models"
	imagetracking "github.com/jackidu14/pholio/internal/services/image-tracking"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type imageTrackingProps struct {
	FakeRecordId      primitive.ObjectID
	NewlyInserted     *mongo.InsertOneResult
	RetreivedInserted *models.RecordImageTracking
}

func TestAddRecordImageTracking(t *testing.T) {
	testData, err := prepareImageTrackingTest()

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, testData.NewlyInserted.InsertedID, testData.RetreivedInserted.Id)
	assert.Equal(t, testData.FakeRecordId, testData.RetreivedInserted.RecordId)
}

func TestUpdateRecordImageTracking(t *testing.T) {
	fakeData, err := prepareImageTrackingTest()

	if err != nil {
		t.Error(err)
		return
	}

	toUpdate := bson.D{
		primitive.E{Key: "Uploaded", Value: true},
		primitive.E{Key: "ThumbConverting", Value: models.ISConverting},
		primitive.E{Key: "MidConverting", Value: models.ISDone},
	}

	updated, err := imagetracking.UpdateRecordImageTracking(fakeData.FakeRecordId.Hex(), toUpdate)
	if err != nil {
		t.Error(err)
		return
	}

	assert.EqualValues(t, updated.ModifiedCount, 1)

	var updatedDocument models.RecordImageTracking
	collection := connector.GetCollection("image-status")
	filter := bson.D{primitive.E{Key: "_id", Value: fakeData.NewlyInserted.InsertedID}}

	err = collection.FindOne(context.TODO(), filter).Decode(&updatedDocument)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, updatedDocument.Uploaded, true)
	assert.Equal(t, updatedDocument.ThumbConverting, models.ISConverting)
	assert.Equal(t, updatedDocument.MidConverting, models.ISDone)
}

func prepareImageTrackingTest() (imageTrackingProps, error) {
	var foundDoc models.RecordImageTracking
	fakeRcId := primitive.NewObjectID()

	added, err := imagetracking.AddRecordImageTracking(fakeRcId.Hex())
	if err != nil {
		return imageTrackingProps{fakeRcId, added, nil}, err
	}

	collection := connector.GetCollection("image-status")
	filter := bson.D{primitive.E{Key: "_id", Value: added.InsertedID}}

	err = collection.FindOne(context.TODO(), filter).Decode(&foundDoc)
	return imageTrackingProps{fakeRcId, added, &foundDoc}, err
}
