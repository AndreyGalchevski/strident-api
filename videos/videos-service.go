package videos

import (
	"context"
	"time"

	"github.com/AndreyGalchevski/strident-api/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var videosCollection *mongo.Collection = db.GetCollection(db.DBClient, "videos")

func getVideos() ([]Video, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := videosCollection.Find(ctx, bson.M{})

	var videos []Video

	if err != nil {
		return videos, err
	}

	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleVideo Video

		err = results.Decode(&singleVideo)

		if err != nil {
			return videos, err
		}

		videos = append(videos, singleVideo)
	}

	return videos, nil
}
