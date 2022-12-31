package auth

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/AndreyGalchevski/strident-api/db"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var TokenMaxAge = 10 * time.Minute

const WRONG_CREDENTIALS_ERROR = "that's not the right password"

func getUsersCollection() *mongo.Collection {
	return db.GetCollection(db.GetDBClient(), "users")
}

func login(credentials Credentials) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user User

	err := getUsersCollection().FindOne(ctx, bson.M{"email": credentials.Email}).Decode(&user)

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

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID string `json:"userID"`
	jwt.RegisteredClaims
}

func GenerateToken(userID string) (string, error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenMaxAge)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(candidate string) (string, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(candidate, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", err
		}
		return "", errors.New("bad token")
	}

	if !token.Valid {
		return "", errors.New("bad token")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user User

	objID, _ := primitive.ObjectIDFromHex(claims.UserID)

	err = getUsersCollection().FindOne(ctx, bson.M{"_id": objID}).Decode(&user)

	if err != nil {
		return "", err
	}

	return user.ID.Hex(), nil
}
