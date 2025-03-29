package databasemodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type Record struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	Description string             `json:"description"`
	Location    string             `json:"location"`
	Date        primitive.DateTime `json:"date"`
}
