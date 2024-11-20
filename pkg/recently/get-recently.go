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
	var records []models.Record

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
