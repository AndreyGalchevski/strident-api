package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Document interface {
	SetID(id primitive.ObjectID)
	GetID() string
}

type Doc struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
}

func (doc *Doc) SetID(id primitive.ObjectID) {
	doc.ID = id
}

func (doc *Doc) GetID() string {
	return doc.ID.Hex()
}

type Collection[T Document] struct {
	collection *mongo.Collection
}

func (repo *Collection[T]) List() ([]T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := repo.collection.Find(ctx, bson.M{})

	var documents []T

	if err != nil {
		return documents, err
	}

	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleDocument T

		err = results.Decode(&singleDocument)

		if err != nil {
			return documents, err
		}

		documents = append(documents, singleDocument)
	}

	return documents, nil
}

func (repo *Collection[T]) Retrieve(id string) (T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var document T

	objID, _ := primitive.ObjectIDFromHex(id)

	err := repo.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&document)

	if err != nil {
		return document, err
	}

	return document, nil
}

func (repo *Collection[T]) Find(filter interface{}) (T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var document T

	err := repo.collection.FindOne(ctx, filter).Decode(&document)

	if err != nil {
		return document, err
	}

	return document, nil
}

func (repo *Collection[T]) Create(params T) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if params.GetID() == "" {
		params.SetID(primitive.NewObjectID())
	}

	result, err := repo.collection.InsertOne(ctx, params)

	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (repo *Collection[T]) Update(id string, params T) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)

	result, err := repo.collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": params})

	if err != nil {
		return false, err
	}

	ok := result.MatchedCount == 1

	return ok, nil
}

func (repo *Collection[T]) Delete(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)

	result, err := repo.collection.DeleteOne(ctx, bson.M{"_id": objID})

	if err != nil {
		return false, err
	}

	ok := result.DeletedCount == 1

	return ok, nil
}
