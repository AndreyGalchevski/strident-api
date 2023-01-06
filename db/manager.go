package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ctx = context.TODO()

func Connect() {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_URI")))

	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	log.Println("Successfully connected to the DB")

	dbClient = client
}

var dbClient *mongo.Client

func GetDBClient() *mongo.Client {
	return dbClient
}

func GetCollection(collectionName string) *mongo.Collection {
	collection := dbClient.Database("main").Collection(collectionName)
	return collection
}
