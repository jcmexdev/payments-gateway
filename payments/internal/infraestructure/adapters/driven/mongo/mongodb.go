package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"payments/config"
)

var mongoClient *mongo.Client
var mongoDatabase *mongo.Database

func ConnectMongoDB(ctx context.Context) {
	clientOptions := options.Client()
	clientOptions.ApplyURI(config.Conf.MongoUrl)
	clientOptions.SetAppName(config.Conf.MongoUrl)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	databaseOptions := options.Database()

	mongoClient = client
	mongoDatabase = mongoClient.Database(config.Conf.MongoDatabase, databaseOptions)
	println("MongoDB is connected")
}

func GetDatabase() *mongo.Database {
	return mongoDatabase
}

func DisconnectMongoDB(ctx context.Context) {
	if mongoClient == nil {
		println("mongodb client is nil")
		return
	}

	err := mongoClient.Disconnect(ctx)
	if err != nil {
		fmt.Println(err)
	}

	println("mongodb is disconnected")
}
