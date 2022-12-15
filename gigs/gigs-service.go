package gigs

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

var gigsCollection *mongo.Collection = db.GetCollection(db.DBClient, "gigs")

func getGigs() ([]Gig, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := gigsCollection.Find(ctx, bson.M{})

	var gigs []Gig

	if err != nil {
		return gigs, err
	}

	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleGig Gig

		err = results.Decode(&singleGig)

		if err != nil {
			return gigs, err
		}

		gigs = append(gigs, singleGig)
	}

	return gigs, nil
}

func getGigByID(id string) (Gig, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var gig Gig

	objID, _ := primitive.ObjectIDFromHex(id)

	err := gigsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&gig)

	if err != nil {
		return gig, err
	}

	return gig, nil
}

func createGig(params GigFormData, image multipart.File) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var gigData Gig
	gigData.ID = primitive.NewObjectID()
	gigData.Name = params.Name
	gigData.Venue = params.Venue
	gigData.Address = params.Address
	gigData.City = params.City
	gigData.Date = params.Date
	gigData.FBEvent = params.FBEvent

	result, err := gigsCollection.InsertOne(ctx, gigData)

	if err != nil {
		return "", err
	}

	imageURL, err := images.UploadImage("gigs", image)

	if err != nil {
		gigsCollection.DeleteOne(ctx, bson.M{"_id": result.InsertedID})
		return "", err
	}

	_, err = gigsCollection.UpdateByID(ctx, result.InsertedID, bson.M{"$set": bson.M{"image": imageURL}})

	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func updateGig(gigID string, params GigFormData, image multipart.File) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(gigID)

	update := bson.M{
		"name":    params.Name,
		"venue":   params.Venue,
		"address": params.Address,
		"city":    params.City,
		"date":    params.Date,
		"fbEvent": params.FBEvent,
	}

	result, err := gigsCollection.UpdateByID(ctx, objID, bson.M{"$set": update})

	if err != nil {
		return false, err
	}

	if result.ModifiedCount != 1 {
		return false, nil
	}

	if image != nil {
		var gig Gig

		gigsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&gig)

		err = images.DeleteImage(gig.Image)

		if err != nil {
			return false, errors.New("failed to delete the old gig image")
		}

		imageURL, err := images.UploadImage("gigs", image)

		if err != nil {
			return false, errors.New("failed to upload the new gig image")
		}

		_, err = gigsCollection.UpdateByID(ctx, objID, bson.M{"$set": bson.M{"image": imageURL}})

		if err != nil {
			return false, err
		}

	}

	return true, nil
}

func deleteGig(gigID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(gigID)
	filter := bson.M{"_id": objID}

	var gigToDelete Gig

	err := gigsCollection.FindOne(ctx, filter).Decode(&gigToDelete)

	if err != nil {
		return false, err
	}

	result, err := gigsCollection.DeleteOne(ctx, filter)

	if err != nil {
		return false, err
	}

	if result.DeletedCount != 1 {
		return false, nil
	}

	err = images.DeleteImage(gigToDelete.Image)

	if err != nil {
		return false, errors.New("failed to delete the gig image")
	}

	return true, nil
}
