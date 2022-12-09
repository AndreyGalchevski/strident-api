package songs

import (
	"context"
	"time"

	"github.com/AndreyGalchevski/strident-api/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var songsCollection *mongo.Collection = db.GetCollection(db.DBClient, "songs")

func getSongs() ([]Song, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := songsCollection.Find(ctx, bson.M{})

	var songs []Song

	if err != nil {
		return songs, err
	}

	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleSong Song

		err = results.Decode(&singleSong)

		if err != nil {
			return songs, err
		}

		songs = append(songs, singleSong)
	}

	return songs, nil
}
