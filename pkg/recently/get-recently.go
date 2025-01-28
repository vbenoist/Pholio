package recently

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackidu14/pholio/internal/database/connector"
	"github.com/jackidu14/pholio/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetRecentlyContent(c *gin.Context) {
	/* Connecting to recently collection */
	collection := connector.GetCollection("recently")
	/* Getting collections. bson.D{} stands for "no filter" */
	cursor, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		fmt.Printf("%s\n", err)
		c.JSON(500, gin.H{"error::database": "Error while reading database - unable to get recently records"})
		return
	}
	defer cursor.Close(context.Background())

	/* Extracting & unmarshall database collections */
	var documents []models.Record
	err = cursor.All(context.Background(), &documents)

	if err != nil {
		fmt.Printf("%s\n", err)
		c.JSON(500, gin.H{"error::database": "Error while reading database - unable to extract recently records"})
		return
	}

	for _, doc := range documents {
		fmt.Printf("Date: %s\n", doc.Date.Time())
	}

	/* Formatting as expected */
	extractedRecords, err := extractRecentlyRecords(documents)

	if err != nil {
		fmt.Printf("%s\n", err)
		c.JSON(500, gin.H{"error::database": "Error while reading data - unable to re-order records"})
		return
	}

	c.JSON(200, extractedRecords)
}

func extractRecentlyRecords(records []models.Record) (models.RecentlyRecords, error) {
	var recentlyRecords models.RecentlyRecords
	lastlyRecord, err := getLastlyRecord(records)

	if err != nil {
		return recentlyRecords, err
	}

	/* 2 weeks before last records */
	dateDelimiter := lastlyRecord.Date.Time().AddDate(0, 0, -14) // 24 * 14

	fmt.Printf("Date to delimit: %s\n", dateDelimiter)

	for _, rec := range records {
		if rec.Date.Time().Before(dateDelimiter) {
			fmt.Printf("Adding %s in lately\n", rec.Date.Time())
			recentlyRecords.Lately = append(recentlyRecords.Lately, rec)
		} else {
			fmt.Printf("Adding %s in lastly\n", rec.Date.Time())
			recentlyRecords.Lastly = append(recentlyRecords.Lastly, rec)
		}
	}

	return recentlyRecords, nil
}

func getLastlyRecord(records []models.Record) (models.Record, error) {
	var result models.Record

	if len(records) == 0 {
		return result, fmt.Errorf("getLastlyRecord: no records to walk in")
	}
	result = records[0]

	for key, rec := range records {
		if key == 0 {
			continue
		}

		if rec.Date.Time().After(result.Date.Time()) {
			result = rec
		}
	}

	return result, nil
}
