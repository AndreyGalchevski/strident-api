package merchandise

import (
	"context"
	"time"

	"github.com/AndreyGalchevski/strident-api/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var merchandiseCollection *mongo.Collection = db.GetCollection(db.DBClient, "merchandise")

func getMerchandise() ([]Merchandise, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := merchandiseCollection.Find(ctx, bson.M{})

	var merchandise []Merchandise

	if err != nil {
		return merchandise, err
	}

	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleMerchandise Merchandise

		err = results.Decode(&singleMerchandise)

		if err != nil {
			return merchandise, err
		}

		merchandise = append(merchandise, singleMerchandise)
	}

	return merchandise, nil
}

func getMerchandiseByID(id string) (Merchandise, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var merchandise Merchandise

	objID, _ := primitive.ObjectIDFromHex(id)

	err := merchandiseCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&merchandise)

	if err != nil {
		return merchandise, err
	}

	return merchandise, nil
}

func createMerchandise(merchandiseData Merchandise) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	merchandiseData.ID = primitive.NewObjectID()

	result, err := merchandiseCollection.InsertOne(ctx, merchandiseData)

	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func updateMerchandise(merchandiseID string, merchandiseData Merchandise) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(merchandiseID)

	update := bson.M{
		"name":  merchandiseData.Name,
		"type":  merchandiseData.Type,
		"price": merchandiseData.Price,
		"url":   merchandiseData.URL,
		"image": merchandiseData.Image,
	}

	result, err := merchandiseCollection.UpdateByID(ctx, objID, bson.M{"$set": update})

	if err != nil {
		return false, err
	}

	ok := result.MatchedCount == 1

	return ok, nil
}
