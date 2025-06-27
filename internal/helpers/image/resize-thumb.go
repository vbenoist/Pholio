package image

import (
	"log"

	databasemodels "github.com/vbenoist/pholio/internal/models/database"

	imagetracking "github.com/vbenoist/pholio/internal/services/image-tracking"
	"github.com/vbenoist/pholio/pkg/helpers/image"

	"github.com/nfnt/resize"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ResizeImageThumb(recordId string, fullPath string) {
	_, err := imagetracking.UpdateRecordImageTracking(recordId,
		bson.D{primitive.E{Key: "ThumbConverting", Value: databasemodels.ISConverting}})

	if err != nil {
		logThumbDatabaseErr(err)
	}

	resizeParams := image.ResizeConfig{
		FullPath:     fullPath,
		ImgWidth:     400,
		ImgType:      image.Thumb,
		CompressAlgo: resize.Lanczos3,
		ForceToJpeg:  true,
	}

	err = image.ResizeImage(resizeParams)

	if err != nil {
		log.Printf("helpers:image::ResizeImageThumb an error occured: %s", err.Error())
		_, err = imagetracking.UpdateRecordImageTracking(recordId,
			bson.D{primitive.E{Key: "ThumbConverting", Value: databasemodels.ISFailed}})

		if err != nil {
			logThumbDatabaseErr(err)
		}
		return
	}

	_, err = imagetracking.UpdateRecordImageTracking(recordId,
		bson.D{primitive.E{Key: "ThumbConverting", Value: databasemodels.ISDone}})

	if err != nil {
		logThumbDatabaseErr(err)
	}
}

func logThumbDatabaseErr(err error) {
	log.Printf("helpers:image::ResizeImageThumb an error occured when updating database status: %s", err.Error())
}
