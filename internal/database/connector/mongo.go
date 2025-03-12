package connector

import (
	"context"
	"fmt"

	"github.com/jackidu14/pholio/internal/helpers/cfg"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
	config := cfg.SetServerConfig()
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.Database.Auth.Username, config.Database.Auth.Password, config.Database.Host, config.Database.Port)
	fmt.Printf("Attempting to connect ... %s\n", uri)
	clientOptions := options.Client().ApplyURI(uri)

	dbclient, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	err = dbclient.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB")

	return dbclient
}

func GetCollection(name string) *mongo.Collection {
	return GetDatabase().Collection(name)
}

func GetDatabase() *mongo.Database {
	config := cfg.SetServerConfig()
	pholioDatabase := client.Database(config.Database.Name)

	if pholioDatabase == nil {
		panic("No matching database has been found.")
	}

	return pholioDatabase
}

func Disconnect() {
	database.Client().Disconnect(context.Background())
	client.Disconnect(context.Background())
}

var database *mongo.Database = GetDatabase()
var client *mongo.Client = Connect()
