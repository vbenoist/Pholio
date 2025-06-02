package record

import (
	"github.com/gin-gonic/gin"
	"github.com/jackidu14/pholio/internal/helpers/controller"
	"github.com/jackidu14/pholio/internal/services/record"
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

	c.JSON(200, paginatedResult)
}
