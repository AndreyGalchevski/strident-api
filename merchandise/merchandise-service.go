package merchandise

import (
	"context"
	"time"

	"github.com/AndreyGalchevski/strident-api/db"
	"go.mongodb.org/mongo-driver/bson"
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