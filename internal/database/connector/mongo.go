package connector

import (
	"context"
	"fmt"

	"github.com/vbenoist/pholio/internal/helpers/cfg"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbInst *mongo.Client

func Connect() *mongo.Client {
	if dbInst != nil {
		fmt.Println("MongoDB: Re-using socket")
		return dbInst
	}

	config := cfg.GetServerConfig()
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.Database.Auth.Username, config.Database.Auth.Password, config.Database.Host, config.Database.Port)
	clientOptions := options.Client().ApplyURI(uri)

	var err error
	fmt.Printf("MongoDB: Attempting to connect ... %s\n", uri)
	dbInst, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	err = dbInst.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("MongoDB: Connected to MongoDB")

	return dbInst
}

func GetCollection(name string) *mongo.Collection {
	return GetDatabase().Collection(name)
}

func GetDatabase() *mongo.Database {
	config := cfg.GetServerConfig()
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
