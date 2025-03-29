package databasemodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type Admin struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Username string             `json:"username"`
	Password string             `json:"password"`
}
