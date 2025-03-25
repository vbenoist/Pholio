package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ImageStatus int

const (
	ISNone ImageStatus = iota
	ISUploading
	ISConverting
	ISFailed
	ISDone
)

type RecordImageTracking struct {
	Id              primitive.ObjectID `bson:"_id" json:"id"`
	RecordId        primitive.ObjectID
	Uploaded        bool
	ThumbConverting ImageStatus
	MidConverting   ImageStatus
}
