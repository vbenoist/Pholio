package record

import (
	"github.com/gin-gonic/gin"
	"github.com/vbenoist/pholio/internal/helpers/controller"
	databasemodels "github.com/vbenoist/pholio/internal/models/database"
	"github.com/vbenoist/pholio/internal/services/record"
)

func GetRecentlyContent(c *gin.Context) {
	paginationParams := controller.GetPaginationParameters(c)
	paginationParams.SortBy = "date"
	paginationParams.SortAsc = -1

	paginatedResult, err := record.GetRecentlyRecords(paginationParams)
	if err != nil {
		c.JSON(500, gin.H{"error::database": "Error while reading database - unable to get recently records"})
		return
	}

	if paginatedResult.Document.Lastly == nil {
		paginatedResult.Document.Lastly = []databasemodels.Record{}
	}
	if paginatedResult.Document.Lately == nil {
		paginatedResult.Document.Lately = []databasemodels.Record{}
	}

	c.JSON(200, paginatedResult)
}
