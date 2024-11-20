package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Record struct {
	Id          primitive.ObjectID `json:"_id,omitempty"`
	NativImgSrc string             `json:"NativImgSrc,omitempty" validate:"required"`
	MidImgSrc   string             `json:"MidImgSrc,omitempty" validate:"required"`
	ThumbImgSrc string             `json:"ThumbImgSrc,omitempty" validate:"required"`
	Description string             `json:"Description,omitempty"`
	Location    string             `json:"Location,omitempty" validate:"required"`
	Date        primitive.DateTime `json:"Date,omitempty" validate:"required"`
}
