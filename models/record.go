package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Record struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	NativImgSrc string             `json:"nativImgSrc"`
	MidImgSrc   string             `json:"midImgSrc"`
	ThumbImgSrc string             `json:"thumbImgSrc"`
	Description string             `json:"description"`
	Location    string             `json:"location"`
	Date        primitive.DateTime `json:"date"`
}
