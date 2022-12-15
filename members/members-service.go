package members

import (
	"context"
	"errors"
	"mime/multipart"
	"time"

	"github.com/AndreyGalchevski/strident-api/db"
	"github.com/AndreyGalchevski/strident-api/images"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var membersCollection *mongo.Collection = db.GetCollection(db.DBClient, "members")

func getMembers() ([]Member, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := membersCollection.Find(ctx, bson.M{})

	var members []Member

	if err != nil {
		return members, err
	}

	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleMember Member

		err = results.Decode(&singleMember)

		if err != nil {
			return members, err
		}

		members = append(members, singleMember)
	}

	return members, nil
}

func getMemberByID(id string) (Member, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var member Member

	objID, _ := primitive.ObjectIDFromHex(id)

	err := membersCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&member)

	if err != nil {
		return member, err
	}

	return member, nil
}

func createMember(params MemberFormData, image multipart.File) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var memberData Member
	memberData.ID = primitive.NewObjectID()
	memberData.Name = params.Name
	memberData.Instrument = params.Instrument

	result, err := membersCollection.InsertOne(ctx, memberData)

	if err != nil {
		return "", err
	}

	imageURL, err := images.UploadImage("members", image)

	if err != nil {
		membersCollection.DeleteOne(ctx, bson.M{"_id": result.InsertedID})
		return "", err
	}

	_, err = membersCollection.UpdateByID(ctx, result.InsertedID, bson.M{"$set": bson.M{"image": imageURL}})

	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func updateMember(memberID string, params MemberFormData, image multipart.File) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(memberID)

	update := bson.M{
		"name":       params.Name,
		"instrument": params.Instrument,
	}

	result, err := membersCollection.UpdateByID(ctx, objID, bson.M{"$set": update})

	if err != nil {
		return false, err
	}

	if result.MatchedCount != 1 {
		return false, nil
	}

	if image != nil {
		var member Member

		membersCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&member)

		err = images.DeleteImage(member.Image)

		if err != nil {
			return false, errors.New("failed to delete the old member image")
		}

		imageURL, err := images.UploadImage("members", image)

		if err != nil {
			return false, errors.New("failed to upload the new member image")
		}

		_, err = membersCollection.UpdateByID(ctx, objID, bson.M{"$set": bson.M{"image": imageURL}})

		if err != nil {
			return false, err
		}

	}

	return true, nil
}

func deleteMember(memberD string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(memberD)
	filter := bson.M{"_id": objID}

	var memberToDelete Member

	err := membersCollection.FindOne(ctx, filter).Decode(&memberToDelete)

	if err != nil {
		return false, err
	}

	_, err = membersCollection.DeleteOne(ctx, filter)

	if err != nil {
		return false, err
	}

	err = images.DeleteImage(memberToDelete.Image)

	if err != nil {
		return false, errors.New("failed to delete the member image")
	}

	return true, nil
}
