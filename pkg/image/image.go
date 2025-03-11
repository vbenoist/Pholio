package image

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackidu14/pholio/internal/database/connector"
	"github.com/jackidu14/pholio/internal/helpers/cfg"
	"github.com/jackidu14/pholio/internal/helpers/image"
	"github.com/jackidu14/pholio/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddImage(c *gin.Context) {
	/* Getting file from form data & checking content type */
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error::file": err.Error()})
		return
	}

	fileType := strings.Split(file.Header.Get("Content-Type"), "/")
	if fileType[0] != "image" || (fileType[1] != "png" && fileType[1] != "jpeg") {
		c.JSON(400, gin.H{"error::file": "Submitted file must be of an image type"})
		return
	}

	/* Checking if given record id exists in database */
	recordId := c.Param("id")
	objRecordId, _ := primitive.ObjectIDFromHex(recordId)
	filter := bson.D{primitive.E{Key: "_id", Value: objRecordId}}

	var result models.Record
	collection := connector.GetCollection("records")

	err = collection.FindOne(c, filter).Decode(&result)
	if err != nil {
		c.JSON(500, gin.H{"error::database": "Error while retreiving related record in database"})
		return
	}

	/* Building final path destination & writing file */
	config := cfg.SetServerConfig()
	fullPath := fmt.Sprintf("%s/%s/original.%s", config.FileManager.UploadPath, recordId, fileType[1])

	err = c.SaveUploadedFile(file, fullPath)
	if err != nil {
		c.JSON(500, gin.H{"error::file": "Error while writing file. Please ensure the configured uploads folder is writable"})
		return
	}

	/* No need to wait for theses tasks */
	go image.ResizeImageThumb(fullPath)
	go image.ResizeImageMid(fullPath)

	c.JSON(200, "Done")
}
