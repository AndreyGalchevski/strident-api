package songs

import (
	"context"
	"time"

	"github.com/AndreyGalchevski/strident-api/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getSongsCollection() *mongo.Collection {
	return db.GetCollection(db.GetDBClient(), "songs")
}

func getSongs() ([]Song, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := getSongsCollection().Find(ctx, bson.M{})

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

	err := getSongsCollection().FindOne(ctx, bson.M{"_id": objID}).Decode(&song)

	if err != nil {
		return song, err
	}

	return song, nil
}

func createSong(params SongFormData) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var songData Song
	songData.ID = primitive.NewObjectID()
	songData.Name = params.Name
	songData.Album = params.Album
	songData.URL = params.URL

	result, err := getSongsCollection().InsertOne(ctx, songData)

	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func updateSong(songID string, params SongFormData) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(songID)

	update := bson.M{
		"name":  params.Name,
		"url":   params.URL,
		"album": params.Album,
	}

	result, err := getSongsCollection().UpdateByID(ctx, objID, bson.M{"$set": update})

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

	result, err := getSongsCollection().DeleteOne(ctx, bson.M{"_id": objID})

	if err != nil {
		return false, err
	}

	ok := result.DeletedCount == 1

	return ok, nil
}
