package image

import (
	"log"

	imagetracking "github.com/jackidu14/pholio/internal/services/image-tracking"
	"github.com/jackidu14/pholio/models"
	"github.com/nfnt/resize"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ResizeImageMid(recordId string, fullPath string) {
	_, err := imagetracking.UpdateRecordImageTracking(recordId,
		bson.D{primitive.E{Key: "MidConverting", Value: models.ISConverting}})

	if err != nil {
		logMidDatabaseErr(err)
	}

	resizeParams := ResizeConfig{
		FullPath:     fullPath,
		ImgWidth:     800,
		ImgType:      Mid,
		CompressAlgo: resize.Lanczos3,
		ForceToJpeg:  false,
	}

	err = resizeImage(resizeParams)

	if err != nil {
		log.Printf("helpers:image::ResizeImageMid an error occured: %s", err.Error())
		_, err = imagetracking.UpdateRecordImageTracking(recordId,
			bson.D{primitive.E{Key: "MidConverting", Value: models.ISFailed}})

		if err != nil {
			logMidDatabaseErr(err)
		}
	}

	_, err = imagetracking.UpdateRecordImageTracking(recordId,
		bson.D{primitive.E{Key: "MidConverting", Value: models.ISDone}})

	if err != nil {
		logMidDatabaseErr(err)
	}
}

func logMidDatabaseErr(err error) {
	log.Printf("helpers:image::ResizeImageMid an error occured when updating database status: %s", err.Error())
}
