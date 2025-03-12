package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DraftRecord struct {
	Description string             `json:"description" binding:"max=100"`
	Location    string             `json:"location" binding:"required,min=4,max=50"`
	Date        primitive.DateTime `json:"date,omitempty" binding:"required"`
}
