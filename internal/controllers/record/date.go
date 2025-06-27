package record

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vbenoist/pholio/internal/helpers/controller"
	"github.com/vbenoist/pholio/internal/services/record"
)

func GetRecordsPerDate(c *gin.Context) {
	paginationParams := controller.GetPaginationParameters(c)
	paginationParams.SortBy = "date"
	paginationParams.SortAsc = -1
	paginationParams.SortAscGroup = -1

	paginatedResult, err := record.GetRecordsGroupByDate(paginationParams)
	// paginatedResult, err := record.GetRecordsGroupByDateDebug(paginationParams)
	if err != nil {
		fmt.Printf("record::GetRecordsPerDate %s\n", err.Error())
		c.JSON(500, gin.H{"error::database": "Error while reading database - unable to get records"})
		return
	}

	c.JSON(200, paginatedResult)
}
