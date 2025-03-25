package image

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackidu14/pholio/internal/helpers/cfg"
	"github.com/jackidu14/pholio/internal/helpers/file"
	imageIntlHelper "github.com/jackidu14/pholio/internal/helpers/image"
	imagetracking "github.com/jackidu14/pholio/internal/services/image-tracking"
	imageGlobHelper "github.com/jackidu14/pholio/pkg/helpers/image"
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

	/* Building final path destination & writing file */
	relatedRecordId := c.Param("id")
	config := cfg.SetServerConfig()
	fullPath := fmt.Sprintf("%s/%s/original.%s", config.FileManager.UploadPath, relatedRecordId, fileType[1])

	err = c.SaveUploadedFile(file, fullPath)
	if err != nil {
		c.JSON(500, gin.H{"error::file": "Error while writing file. Please ensure the configured uploads folder is writable"})
		return
	}

	/* Updating image tracking status */
	_, err = imagetracking.UpdateRecordImageTracking(relatedRecordId, bson.D{primitive.E{Key: "Uploaded", Value: true}})

	if err != nil {
		c.JSON(500, "Err")
		return
	}

	/* No need to wait for theses tasks */
	go imageIntlHelper.ResizeImageThumb(relatedRecordId, fullPath)
	go imageIntlHelper.ResizeImageMid(relatedRecordId, fullPath)

	c.JSON(200, "Done")
}

func GetThumbImage(c *gin.Context) {
	fullPath, err := file.GetFileFullpath(c.Param("id"), imageGlobHelper.Thumb)

	if err != nil {
		c.JSON(400, gin.H{"error::file": "No thumb found for this record"})
		return
	}

	http.ServeFile(c.Writer, c.Request, fullPath)
}

func GetMidImage(c *gin.Context) {
	fullPath, err := file.GetFileFullpath(c.Param("id"), imageGlobHelper.Mid)

	if err != nil {
		c.JSON(400, gin.H{"error::file": "No midsize found for this record"})
		return
	}

	http.ServeFile(c.Writer, c.Request, fullPath)
}

func GetOrigImage(c *gin.Context) {
	fullPath, err := file.GetFileFullpath(c.Param("id"), imageGlobHelper.Orig)

	if err != nil {
		c.JSON(400, gin.H{"error::file": "No image found for this record"})
		return
	}

	http.ServeFile(c.Writer, c.Request, fullPath)
}
