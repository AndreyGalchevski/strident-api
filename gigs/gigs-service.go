package gigs

import (
	"context"
	"errors"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/AndreyGalchevski/strident-api/db"
	"github.com/AndreyGalchevski/strident-api/images"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getGigsCollection() *mongo.Collection {
	return db.GetCollection(db.GetDBClient(), "gigs")
}

func getGigs() ([]Gig, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Find().SetSort(bson.D{{Key: "date", Value: -1}})

	results, err := getGigsCollection().Find(ctx, bson.M{}, opts)

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

	err := getGigsCollection().FindOne(ctx, bson.M{"_id": objID}).Decode(&gig)

	if err != nil {
		return gig, err
	}

	return gig, nil
}

func createGig(params GigFormData, image multipart.File) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	timestamp, err := strconv.ParseInt(params.Date, 10, 0)

	if err != nil {
		return "", err
	}

	var gigData Gig
	gigData.ID = primitive.NewObjectID()
	gigData.Name = params.Name
	gigData.Venue = params.Venue
	gigData.Address = params.Address
	gigData.City = params.City
	gigData.Date = primitive.NewDateTimeFromTime(time.UnixMilli(timestamp))
	gigData.FBEvent = params.FBEvent

	result, err := getGigsCollection().InsertOne(ctx, gigData)

	if err != nil {
		return "", err
	}

	imageURL, err := images.UploadImage("gigs", image)

	if err != nil {
		getGigsCollection().DeleteOne(ctx, bson.M{"_id": result.InsertedID})
		return "", err
	}

	_, err = getGigsCollection().UpdateByID(ctx, result.InsertedID, bson.M{"$set": bson.M{"image": imageURL}})

	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func updateGig(gigID string, params GigFormData, image multipart.File) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(gigID)

	timestamp, err := strconv.ParseInt(params.Date, 10, 0)

	if err != nil {
		return false, err
	}
	update := bson.M{
		"name":    params.Name,
		"venue":   params.Venue,
		"address": params.Address,
		"city":    params.City,
		"date":    primitive.NewDateTimeFromTime(time.UnixMilli(timestamp)),
		"fbEvent": params.FBEvent,
	}

	result, err := getGigsCollection().UpdateByID(ctx, objID, bson.M{"$set": update})

	if err != nil {
		return false, err
	}

	if result.MatchedCount != 1 {
		return false, nil
	}

	if image != nil {
		var gig Gig

		getGigsCollection().FindOne(ctx, bson.M{"_id": objID}).Decode(&gig)

		err = images.DeleteImage(gig.Image)

		if err != nil {
			return false, errors.New("failed to delete the old gig image")
		}

		imageURL, err := images.UploadImage("gigs", image)

		if err != nil {
			return false, errors.New("failed to upload the new gig image")
		}

		_, err = getGigsCollection().UpdateByID(ctx, objID, bson.M{"$set": bson.M{"image": imageURL}})

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

	err := getGigsCollection().FindOne(ctx, filter).Decode(&gigToDelete)

	if err != nil {
		return false, err
	}

	result, err := getGigsCollection().DeleteOne(ctx, filter)

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
