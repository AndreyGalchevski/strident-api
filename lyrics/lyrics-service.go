package lyrics

import (
	"context"
	"time"

	"github.com/AndreyGalchevski/strident-api/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func getLyricByID(id string) (Lyric, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var lyric Lyric

	objID, _ := primitive.ObjectIDFromHex(id)

	err := lyricsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&lyric)

	if err != nil {
		return lyric, err
	}

	return lyric, nil
}

func createLyric(lyricData Lyric) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	lyricData.ID = primitive.NewObjectID()

	result, err := lyricsCollection.InsertOne(ctx, lyricData)

	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func updateLyric(lyricID string, lyricData Lyric) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(lyricID)

	update := bson.M{
		"name":  lyricData.Name,
		"text":  lyricData.Text,
		"album": lyricData.Album,
	}

	result, err := lyricsCollection.UpdateByID(ctx, objID, bson.M{"$set": update})

	if err != nil {
		return false, err
	}

	ok := result.MatchedCount == 1

	return ok, nil
}
