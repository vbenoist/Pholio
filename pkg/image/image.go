package image

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackidu14/pholio/internal/helpers/cfg"
	"github.com/jackidu14/pholio/internal/helpers/file"
	"github.com/jackidu14/pholio/internal/helpers/image"
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
	config := cfg.SetServerConfig()
	fullPath := fmt.Sprintf("%s/%s/original.%s", config.FileManager.UploadPath, c.Param("id"), fileType[1])

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

func GetThumbImage(c *gin.Context) {
	fullPath := file.GetFileFullpath(c.Param("id"), image.Thumb)
	http.ServeFile(c.Writer, c.Request, fullPath)
}

func GetMidImage(c *gin.Context) {
	fullPath := file.GetFileFullpath(c.Param("id"), image.Mid)
	http.ServeFile(c.Writer, c.Request, fullPath)
}

func GetOrigImage(c *gin.Context) {
	fullPath := file.GetFileFullpath(c.Param("id"), image.Orig)
	http.ServeFile(c.Writer, c.Request, fullPath)
}
