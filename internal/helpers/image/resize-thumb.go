package image

import (
	"log"

	imagetracking "github.com/jackidu14/pholio/internal/services/image-tracking"
	"github.com/jackidu14/pholio/models"
	"github.com/nfnt/resize"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ResizeImageThumb(recordId string, fullPath string) {
	_, err := imagetracking.UpdateRecordImageTracking(recordId,
		bson.D{primitive.E{Key: "ThumbConverting", Value: models.ISConverting}})

	if err != nil {
		logThumbDatabaseErr(err)
	}

	resizeParams := ResizeConfig{
		FullPath:     fullPath,
		ImgWidth:     400,
		ImgType:      Thumb,
		CompressAlgo: resize.Lanczos3,
		ForceToJpeg:  true,
	}

	err = resizeImage(resizeParams)

	if err != nil {
		log.Printf("helpers:image::ResizeImageThumb an error occured: %s", err.Error())
		_, err = imagetracking.UpdateRecordImageTracking(recordId,
			bson.D{primitive.E{Key: "ThumbConverting", Value: models.ISFailed}})

		if err != nil {
			logThumbDatabaseErr(err)
		}
		return
	}

	_, err = imagetracking.UpdateRecordImageTracking(recordId,
		bson.D{primitive.E{Key: "ThumbConverting", Value: models.ISDone}})

	if err != nil {
		logThumbDatabaseErr(err)
	}
}

func logThumbDatabaseErr(err error) {
	log.Printf("helpers:image::ResizeImageThumb an error occured when updating database status: %s", err.Error())
}
