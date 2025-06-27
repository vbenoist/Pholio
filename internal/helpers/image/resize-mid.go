package image

import (
	"log"

	"github.com/nfnt/resize"
	databasemodels "github.com/vbenoist/pholio/internal/models/database"
	imagetracking "github.com/vbenoist/pholio/internal/services/image-tracking"
	"github.com/vbenoist/pholio/pkg/helpers/image"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ResizeImageMid(recordId string, fullPath string) {
	_, err := imagetracking.UpdateRecordImageTracking(recordId,
		bson.D{primitive.E{Key: "MidConverting", Value: databasemodels.ISConverting}})

	if err != nil {
		logMidDatabaseErr(err)
	}

	resizeParams := image.ResizeConfig{
		FullPath:     fullPath,
		ImgWidth:     800,
		ImgType:      image.Mid,
		CompressAlgo: resize.Lanczos3,
		ForceToJpeg:  false,
	}

	err = image.ResizeImage(resizeParams)

	if err != nil {
		log.Printf("helpers:image::ResizeImageMid an error occured: %s", err.Error())
		_, err = imagetracking.UpdateRecordImageTracking(recordId,
			bson.D{primitive.E{Key: "MidConverting", Value: databasemodels.ISFailed}})

		if err != nil {
			logMidDatabaseErr(err)
		}
	}

	_, err = imagetracking.UpdateRecordImageTracking(recordId,
		bson.D{primitive.E{Key: "MidConverting", Value: databasemodels.ISDone}})

	if err != nil {
		logMidDatabaseErr(err)
	}
}

func logMidDatabaseErr(err error) {
	log.Printf("helpers:image::ResizeImageMid an error occured when updating database status: %s", err.Error())
}
