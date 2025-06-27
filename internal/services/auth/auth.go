package auth

import (
	"context"

	"github.com/vbenoist/pholio/internal/database/connector"
	databasemodels "github.com/vbenoist/pholio/internal/models/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func GetAdminFromUid(userId string) (databasemodels.Admin, error) {
	var registeredAdmin databasemodels.Admin
	uObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return registeredAdmin, err
	}

	collection := connector.GetCollection("admin")
	filter := bson.D{
		primitive.E{Key: "_id", Value: uObjId},
	}

	err = collection.FindOne(context.TODO(), filter).Decode(&registeredAdmin)
	return registeredAdmin, err
}

func GetAdminFromIds(username string, password string) (databasemodels.Admin, error) {
	collection := connector.GetCollection("admin")
	filter := bson.D{
		primitive.E{Key: "username", Value: username},
		primitive.E{Key: "password", Value: password}, // add hash
	}

	var registeredAdmin databasemodels.Admin
	err := collection.FindOne(context.TODO(), filter).Decode(&registeredAdmin)
	return registeredAdmin, err
}

func GetAdminFromUsername(username string) (databasemodels.Admin, error) {
	collection := connector.GetCollection("admin")
	filter := bson.D{
		primitive.E{Key: "username", Value: username},
	}

	var registeredAdmin databasemodels.Admin
	err := collection.FindOne(context.TODO(), filter).Decode(&registeredAdmin)
	return registeredAdmin, err
}

func AdminDatabaseExists() (bool, error) {
	collection := connector.GetCollection("admin")
	docsCount, err := collection.EstimatedDocumentCount(context.TODO())
	if err != nil {
		return false, err
	}

	return docsCount > 0, nil
}

func InitAdminDatabase(username string, password string) error {
	collection := connector.GetCollection("admin")

	passwordHash, err := GetPasswordHash(password)
	if err != nil {
		return err
	}

	defaultAdmin := databasemodels.Admin{
		Id:       primitive.NewObjectID(),
		Username: username,
		Password: string(passwordHash),
	}
	_, err = collection.InsertOne(context.TODO(), defaultAdmin)
	return err
}

func GetPasswordHash(pass string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}
