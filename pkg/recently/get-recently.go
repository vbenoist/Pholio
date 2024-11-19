package recently

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackidu14/pholio/internal/database/connector"
	"go.mongodb.org/mongo-driver/bson"
)

// Difference between Lastly / Lately is just from configured date -- TODO
type RecentlyRecords struct {
	Lastly []Record
	Lately []Record
}

type Record struct {
	NativImgSrc string
	MidImgSrc   string
	ThumbImgSrc string
	Description string
	Location    string
	Date        string
}

func GetRecentlyContent(c *gin.Context) {
	var records []Record

	collection := connector.GetCollection("recently")
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		fmt.Printf("%s\n", err)
		c.JSON(500, gin.H{"error::database": "Error while reading database - unable to get recently records"})
		return
	}
	defer cursor.Close(context.Background())

	cursor.All(context.Background(), &records)
	c.JSON(200, records)
}
