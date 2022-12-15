package gigs

import (
	"context"
	"errors"
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

func createGig(gigData Gig) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	gigData.ID = primitive.NewObjectID()

	result, err := gigsCollection.InsertOne(ctx, gigData)

	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func updateGig(gigID string, gigData Gig) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(gigID)

	update := bson.M{
		"name":    gigData.Name,
		"venue":   gigData.Venue,
		"address": gigData.Address,
		"city":    gigData.City,
		"date":    gigData.Date,
		"fbEvent": gigData.FBEvent,
		"image":   gigData.Image,
	}

	result, err := gigsCollection.UpdateByID(ctx, objID, bson.M{"$set": update})

	if err != nil {
		return false, err
	}

	ok := result.MatchedCount == 1

	return ok, nil
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

	_, err = gigsCollection.DeleteOne(ctx, filter)

	if err != nil {
		return false, err
	}

	err = images.DeleteImage(gigToDelete.Image)

	if err != nil {
		return false, errors.New("unable to delete the gig image")
	}

	return true, nil
}
