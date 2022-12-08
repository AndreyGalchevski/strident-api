package auth

import (
	"context"
	"errors"
	"time"

	"github.com/AndreyGalchevski/strident-api/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const WRONG_CREDENTIALS_ERROR = "wrong credentials"

var usersCollection *mongo.Collection = db.GetCollection(db.DBClient, "users")

func login(credentials Credentials) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User

	err := usersCollection.FindOne(ctx, bson.M{"username": credentials.Username}).Decode(&user)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return "", errors.New(WRONG_CREDENTIALS_ERROR)
		}

		return "", err
	}

	if credentials.Password != user.Password {
		return "", errors.New(WRONG_CREDENTIALS_ERROR)
	}

	token, err := GenerateToken(credentials.Username)

	if err != nil {
		return "", err
	}

	return token, nil
}
