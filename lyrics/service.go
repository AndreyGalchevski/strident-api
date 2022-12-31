package lyrics

import (
	"context"
	"time"

	"github.com/AndreyGalchevski/strident-api/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getLyricsCollection() *mongo.Collection {
	return db.GetCollection(db.GetDBClient(), "lyrics")
}

func getLyrics() ([]Lyric, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := getLyricsCollection().Find(ctx, bson.M{})

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

	err := getLyricsCollection().FindOne(ctx, bson.M{"_id": objID}).Decode(&lyric)

	if err != nil {
		return lyric, err
	}

	return lyric, nil
}

func createLyric(params LyricFormData) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var lyricData Lyric
	lyricData.ID = primitive.NewObjectID()
	lyricData.Name = params.Name
	lyricData.Text = params.Text

	result, err := getLyricsCollection().InsertOne(ctx, lyricData)

	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func updateLyric(lyricID string, params LyricFormData) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(lyricID)

	update := bson.M{
		"name": params.Name,
		"text": params.Text,
	}

	result, err := getLyricsCollection().UpdateByID(ctx, objID, bson.M{"$set": update})

	if err != nil {
		return false, err
	}

	ok := result.MatchedCount == 1

	return ok, nil
}

func deleteLyric(lyricID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(lyricID)

	result, err := getLyricsCollection().DeleteOne(ctx, bson.M{"_id": objID})

	if err != nil {
		return false, err
	}

	ok := result.DeletedCount == 1

	return ok, nil
}
