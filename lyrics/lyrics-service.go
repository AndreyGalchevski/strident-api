package lyrics

import (
	"context"
	"time"

	"github.com/AndreyGalchevski/strident-api/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var lyricsCollection *mongo.Collection = db.GetCollection(db.DBClient, "lyrics")

func getLyrics() ([]Lyric, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := lyricsCollection.Find(ctx, bson.M{})

	var lyrics []Lyric

	if err != nil {
		return lyrics, err
	}

	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleLyric Lyric

		err = results.Decode(&singleLyric)

		if err != nil {
			return lyrics, err
		}

		lyrics = append(lyrics, singleLyric)
	}

	return lyrics, nil
}
