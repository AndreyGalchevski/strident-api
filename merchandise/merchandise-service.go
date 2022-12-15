package merchandise

import (
	"context"
	"errors"
	"mime/multipart"
	"time"

	"github.com/AndreyGalchevski/strident-api/db"
	"github.com/AndreyGalchevski/strident-api/images"
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

func createMerchandise(params MerchandiseFormData, image multipart.File) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var merchandiseData Merchandise
	merchandiseData.ID = primitive.NewObjectID()
	merchandiseData.Name = params.Name
	merchandiseData.Type = params.Type
	merchandiseData.Price = params.Price
	merchandiseData.URL = params.URL

	result, err := merchandiseCollection.InsertOne(ctx, merchandiseData)

	if err != nil {
		return "", err
	}

	imageURL, err := images.UploadImage("merchandise", image)

	if err != nil {
		merchandiseCollection.DeleteOne(ctx, bson.M{"_id": result.InsertedID})
		return "", err
	}

	_, err = merchandiseCollection.UpdateByID(ctx, result.InsertedID, bson.M{"$set": bson.M{"image": imageURL}})

	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func updateMerchandise(merchandiseID string, params MerchandiseFormData, image multipart.File) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(merchandiseID)

	update := bson.M{
		"name":  params.Name,
		"type":  params.Type,
		"price": params.Price,
		"url":   params.URL,
	}

	result, err := merchandiseCollection.UpdateByID(ctx, objID, bson.M{"$set": update})

	if err != nil {
		return false, err
	}

	if result.ModifiedCount != 1 {
		return false, nil
	}

	if image != nil {
		var merchandise Merchandise

		merchandiseCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&merchandise)

		err = images.DeleteImage(merchandise.Image)

		if err != nil {
			return false, errors.New("failed to delete the old merchandise image")
		}

		imageURL, err := images.UploadImage("merchandise", image)

		if err != nil {
			return false, errors.New("failed to upload the new merchandise image")
		}

		_, err = merchandiseCollection.UpdateByID(ctx, objID, bson.M{"$set": bson.M{"image": imageURL}})

		if err != nil {
			return false, err
		}

	}

	return true, nil
}

func deleteMerchandise(merchandiseD string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(merchandiseD)
	filter := bson.M{"_id": objID}

	var merchandiseToDelete Merchandise

	err := merchandiseCollection.FindOne(ctx, filter).Decode(&merchandiseToDelete)

	if err != nil {
		return false, err
	}

	result, err := merchandiseCollection.DeleteOne(ctx, filter)

	if err != nil {
		return false, err
	}

	if result.DeletedCount != 1 {
		return false, nil
	}

	err = images.DeleteImage(merchandiseToDelete.Image)

	if err != nil {
		return false, errors.New("failed to delete the merchandise image")
	}

	return true, nil
}
