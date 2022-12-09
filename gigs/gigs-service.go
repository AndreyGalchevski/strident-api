package gigs

import (
	"context"
	"time"

	"github.com/AndreyGalchevski/strident-api/db"
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
