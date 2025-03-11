package record

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackidu14/pholio/internal/database/connector"
	"github.com/jackidu14/pholio/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddRecords(c *gin.Context) {
	var records []models.DraftRecord
	err := c.BindJSON(&records)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error::query": getGenericError(err)})
		return
	}

	collection := connector.GetCollection("records")
	fmt.Printf("Gave: %s - %s - %s\n", records[0].Description, records[0].Location, records[0].Date.Time())

	/*
		MEMO: interface{} = any : https://stackoverflow.com/questions/23148812/whats-the-meaning-of-interface
		_, err = collection.InsertMany(c, []interface{}{records})
	*/
	_, err = collection.InsertMany(c, []any{records})

	if err != nil {
		fmt.Printf("%s\n", err)
		c.JSON(500, gin.H{"error::database": "Error while inserting records in database"})
		return
	}

	c.JSON(200, "Done")
}

func AddRecord(c *gin.Context) {
	var record models.DraftRecord
	err := c.BindJSON(&record)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error::query": getGenericError(err)})
		return
	}

	collection := connector.GetCollection("records")
	res, err := collection.InsertOne(c, record)

	if err != nil {
		fmt.Printf("%s\n", err)
		c.JSON(500, gin.H{"error::database": "Error while inserting records in database"})
		return
	}

	c.JSON(200, res.InsertedID)
}

func EditRecord(c *gin.Context) {
	var record models.DraftRecord
	err := c.BindJSON(&record)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error::query": getGenericError(err)})
		return
	}

	collection := connector.GetCollection("records")
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	res, err := collection.ReplaceOne(c, filter, record)

	if err != nil {
		fmt.Printf("%s\n", err)
		c.JSON(500, gin.H{"error::database": "Error while updating record in database"})
		return
	}

	if res.ModifiedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error::database": "Error while retreiving record in database"})
		return
	}

	c.JSON(http.StatusAccepted, id)
}

func RemoveRecord(c *gin.Context) {
	collection := connector.GetCollection("records")
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	res, err := collection.DeleteOne(c, filter)

	if err != nil {
		fmt.Printf("%s\n", err)
		c.JSON(500, gin.H{"error::database": "Error while deleting record from database"})
		return
	}

	if res.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error::database": "Error while retreiving record to delete from database"})
		return
	}

	c.JSON(http.StatusAccepted, id)
}

func getGenericError(err error) string {
	return fmt.Sprintf("Error while extracting body request content: %s", err.Error())
}

/*
Pouvoir charger plusieurs images, et pour chaque spécifier la description, date, ...
A la validation du formulaire, on arrive ici, avec un tableau de records : créer un
job pour chaque image envoyée -> enregistrement en bdd des infos de base, traitement
dans les tailles intermédiaires, et enregistrement en bdd des tailles intermédiaires

*/
