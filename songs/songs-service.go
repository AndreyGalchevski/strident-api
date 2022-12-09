package songs

import (
	"context"
	"time"

	"github.com/AndreyGalchevski/strident-api/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func getSongByID(id string) (Song, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var song Song

	objID, _ := primitive.ObjectIDFromHex(id)

	err := songsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&song)

	if err != nil {
		return song, err
	}

	return song, nil
}
