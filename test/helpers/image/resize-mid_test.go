package image_test

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/jackidu14/pholio/internal/database/connector"
	"github.com/jackidu14/pholio/internal/helpers/cfg"
	"github.com/jackidu14/pholio/internal/helpers/image"
	"github.com/jackidu14/pholio/internal/models"
	imagetracking "github.com/jackidu14/pholio/internal/services/image-tracking"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestResizeImageMid(t *testing.T) {
	/* Preparing tests */
	fakeRcId, fakeRcFullPath, err := prepareResizeImageTest()
	if err != nil {
		t.Error(err)
		return
	}

	added, err := imagetracking.AddRecordImageTracking(fakeRcId)
	if err != nil {
		t.Error(err)
		return
	}

	/* Checking file has been created */
	image.ResizeImageMid(fakeRcId, fakeRcFullPath)
	fakeMidsizeFilePath := strings.Replace(fakeRcFullPath, "original", "midsize", 1)
	assert.FileExists(t, fakeMidsizeFilePath)

	/* Checking file has been writted */
	midsizeFile, err := os.Open(fakeMidsizeFilePath)
	if err != nil {
		t.Error(err)
		return
	}
	defer midsizeFile.Close()

	midsizeFileInfos, err := midsizeFile.Stat()
	if err != nil {
		t.Error(err)
		return
	}
	assert.True(t, midsizeFileInfos.Size() > 1000)

	/* Checking if tracking has been updated in database */
	var recordImageTracking models.RecordImageTracking
	collection := connector.GetCollection("image-status")
	filter := bson.D{primitive.E{Key: "_id", Value: added.InsertedID}}

	err = collection.FindOne(context.TODO(), filter).Decode(&recordImageTracking)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, recordImageTracking.MidConverting, models.ISDone)
}

func prepareResizeImageTest() (string, string, error) {
	/* Opening base example file */
	exampleFile, err := os.Open("../../static/mock-assets/example.jpg")
	if err != nil {
		return "", "", err
	}
	defer exampleFile.Close()

	config := cfg.SetServerConfig()

	/* Creating a copy of base example file, in future working dir */
	fakeRcId := primitive.NewObjectID()
	// fileExt := strings.Split(exampleFile.Name(), "example.")[1]
	fakeRcFolder := fmt.Sprintf("%s/%s/", config.FileManager.UploadPath, fakeRcId.Hex())
	fakeRcFullPath := fmt.Sprintf("%s/%s/original.%s", config.FileManager.UploadPath, fakeRcId.Hex(), "jpg")

	if err = os.MkdirAll(fakeRcFolder, 0774); err != nil {
		return fakeRcId.Hex(), fakeRcFullPath, err
	}

	out, err := os.Create(fakeRcFullPath)
	if err != nil {
		return fakeRcId.Hex(), fakeRcFullPath, err
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, exampleFile); err != nil {
		return fakeRcId.Hex(), fakeRcFullPath, err
	}

	err = out.Sync()
	return fakeRcId.Hex(), fakeRcFullPath, err
}
