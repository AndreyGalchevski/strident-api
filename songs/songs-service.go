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

func createSong(songData Song) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	songData.ID = primitive.NewObjectID()

	result, err := songsCollection.InsertOne(ctx, songData)

	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func updateSong(songID string, songData Song) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(songID)

	update := bson.M{
		"name":  songData.Name,
		"url":   songData.URL,
		"album": songData.Album,
	}

	result, err := songsCollection.UpdateByID(ctx, objID, bson.M{"$set": update})

	if err != nil {
		return false, err
	}

	ok := result.MatchedCount == 1

	return ok, nil
}

func deleteSong(songID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(songID)

	result, err := songsCollection.DeleteOne(ctx, bson.M{"_id": objID})

	if err != nil {
		return false, err
	}

	ok := result.DeletedCount == 1

	return ok, nil
}
