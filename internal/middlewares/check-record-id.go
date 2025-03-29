package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jackidu14/pholio/internal/database/connector"
	databasemodels "github.com/jackidu14/pholio/internal/models/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckRecordId() gin.HandlerFunc {
	return func(c *gin.Context) {
		recordId := c.Param("id")
		objRecordId, _ := primitive.ObjectIDFromHex(recordId)
		filter := bson.D{primitive.E{Key: "_id", Value: objRecordId}}

		var result databasemodels.Record
		collection := connector.GetCollection("records")

		err := collection.FindOne(c, filter).Decode(&result)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error::check": "No record found"})
			return
		}

		c.Next()
	}
}
