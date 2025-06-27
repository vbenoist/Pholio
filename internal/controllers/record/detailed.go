package record

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vbenoist/pholio/internal/helpers/controller"
	"github.com/vbenoist/pholio/internal/services/record"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDetailedRecords(c *gin.Context) {
	paginationParams := controller.GetPaginationParameters(c)
	results, err := record.GetDetailedRecords(paginationParams)
	if err != nil {
		fmt.Printf("record::GetDetailedRecords %s\n", err.Error())

		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error::query": "No record found"})
			return
		}

		// if err == os.ErrNotExist || err == os.ErrInvalid || os.IsNotExist(err) {
		if os.IsNotExist(err) {
			c.JSON(http.StatusInternalServerError, gin.H{"error::query": "An orphan record has been found: %s"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error::query": "An error occured on record retreiving"})
		return
	}

	c.JSON(http.StatusAccepted, results)
}
