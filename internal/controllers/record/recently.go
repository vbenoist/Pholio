package record

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackidu14/pholio/internal/helpers/controller"
	apimodels "github.com/jackidu14/pholio/internal/models/api"
	databasemodels "github.com/jackidu14/pholio/internal/models/database"
	"github.com/jackidu14/pholio/internal/services/record"
)

func GetRecentlyContent(c *gin.Context) {
	paginationParams := controller.GetPaginationParameters(c)
	paginatedResult, err := record.GetRecords(paginationParams)
	if err != nil {
		c.JSON(500, gin.H{"error::database": "Error while reading database - unable to get recently records"})
		return
	}

	if len(paginatedResult.Documents) == 0 {
		c.JSON(200, paginatedResult)
		return
	}

	/* Formatting as expected */
	extractedRecords, err := extractRecentlyRecords(paginatedResult.Documents)

	if err != nil {
		fmt.Printf("%s\n", err)
		c.JSON(500, gin.H{"error::database": "Error while reading data - unable to re-order records"})
		return
	}

	results := apimodels.PaginatedResult[apimodels.RecentlyRecords]{
		Pagination: paginatedResult.Pagination,
		Document:   *extractedRecords,
	}
	c.JSON(200, results)
}

func extractRecentlyRecords(records []databasemodels.Record) (*apimodels.RecentlyRecords, error) {
	lastlyRecord, err := getLastlyRecord(records)

	if err != nil {
		return nil, err
	}

	/* 2 weeks before last records */
	dateDelimiter := lastlyRecord.Date.Time().AddDate(0, 0, -14) // 24 * 14

	fmt.Printf("Date to delimit: %s\n", dateDelimiter)
	var recentlyRecords apimodels.RecentlyRecords

	for _, rec := range records {
		if rec.Date.Time().Before(dateDelimiter) {
			fmt.Printf("Adding %s in lately\n", rec.Date.Time())
			recentlyRecords.Lately = append(recentlyRecords.Lately, rec)
		} else {
			fmt.Printf("Adding %s in lastly\n", rec.Date.Time())
			recentlyRecords.Lastly = append(recentlyRecords.Lastly, rec)
		}
	}

	return &recentlyRecords, nil
}

func getLastlyRecord(records []databasemodels.Record) (*databasemodels.Record, error) {
	if len(records) == 0 {
		return nil, fmt.Errorf("getLastlyRecord: no records to walk in")
	}
	result := records[0]

	for key, rec := range records {
		if key == 0 {
			continue
		}

		if rec.Date.Time().After(result.Date.Time()) {
			result = rec
		}
	}

	return &result, nil
}
