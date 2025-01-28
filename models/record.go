package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Record struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	NativImgSrc string             `json:"nativImgSrc,omitempty" validate:"required"`
	MidImgSrc   string             `json:"midImgSrc,omitempty" validate:"required"`
	ThumbImgSrc string             `json:"thumbImgSrc,omitempty" validate:"required"`
	Description string             `json:"description,omitempty"`
	Location    string             `json:"location,omitempty" validate:"required"`
	Date        primitive.DateTime `json:"date,omitempty" validate:"required"`
}
