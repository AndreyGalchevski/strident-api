package db

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ctx = context.TODO()

func Connect() *mongo.Client {
	if err := godotenv.Load(); err != nil {
		panic("Can't load env vars for DB connection")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_URI")))

	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged.")

	return client
}

var DB *mongo.Client = Connect()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("main").Collection(collectionName)
	return collection
}
