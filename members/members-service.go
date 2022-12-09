package members

import (
	"context"
	"time"

	"github.com/AndreyGalchevski/strident-api/db"
	"go.mongodb.org/mongo-driver/bson"
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
