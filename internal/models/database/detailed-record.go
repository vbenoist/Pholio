package databasemodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type DetailedRecord struct {
	Record
	Pin []struct {
		PinId primitive.ObjectID `bson:"pin_id"`
	} `bson:"record_pin"`
}
