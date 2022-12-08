package auth

import (
	"context"
	"errors"
	"time"

	"github.com/AndreyGalchevski/strident-api/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const WRONG_CREDENTIALS_ERROR = "that's not the right password"

var usersCollection *mongo.Collection = db.GetCollection(db.DBClient, "users")

func login(credentials Credentials) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User

	err := usersCollection.FindOne(ctx, bson.M{"email": credentials.Email}).Decode(&user)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return "", errors.New(WRONG_CREDENTIALS_ERROR)
		}

		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))

	if err != nil {
		return "", errors.New(WRONG_CREDENTIALS_ERROR)
	}

	token, err := GenerateToken(user.ID.Hex())

	if err != nil {
		return "", err
	}

	return token, nil
}
