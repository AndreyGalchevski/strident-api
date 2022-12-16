package videos

import (
	"context"
	"time"

	"github.com/AndreyGalchevski/strident-api/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getVideosCollection() *mongo.Collection {
	return db.GetCollection(db.GetDBClient(), "videos")
}

func getVideos() ([]Video, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := getVideosCollection().Find(ctx, bson.M{})

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

func getVideoByID(id string) (Video, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var video Video

	objID, _ := primitive.ObjectIDFromHex(id)

	err := getVideosCollection().FindOne(ctx, bson.M{"_id": objID}).Decode(&video)

	if err != nil {
		return video, err
	}

	return video, nil
}

func createVideo(params VideoFormData) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var videoData Video
	videoData.ID = primitive.NewObjectID()
	videoData.Name = params.Name
	videoData.URL = params.URL

	result, err := getVideosCollection().InsertOne(ctx, videoData)

	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func updateVideo(videoID string, params VideoFormData) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(videoID)

	update := bson.M{
		"name": params.Name,
		"url":  params.URL,
	}

	result, err := getVideosCollection().UpdateByID(ctx, objID, bson.M{"$set": update})

	if err != nil {
		return false, err
	}

	ok := result.MatchedCount == 1

	return ok, nil
}

func deleteVideo(videoID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(videoID)

	result, err := getVideosCollection().DeleteOne(ctx, bson.M{"_id": objID})

	if err != nil {
		return false, err
	}

	ok := result.DeletedCount == 1

	return ok, nil
}
