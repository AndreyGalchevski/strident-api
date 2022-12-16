package db

import (
	"context"
	"fmt"
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

	fmt.Println("Successfully connected and pinged.")

	dbClient = client
}

var dbClient *mongo.Client

func GetDBClient() *mongo.Client {
	return dbClient
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("main").Collection(collectionName)
	return collection
}
