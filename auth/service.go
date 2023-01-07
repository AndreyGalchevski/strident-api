package auth

import (
	"errors"
	"os"
	"time"

	"github.com/AndreyGalchevski/strident-api/db"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var TokenMaxAge = 120 * time.Minute

const WRONG_CREDENTIALS_ERROR = "that's not the right password"

func login(credentials Credentials) (string, error) {
	user, err := db.GetDB().Users.Find(bson.M{"email": credentials.Email})

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

	user, err := db.GetDB().Users.Retrieve(claims.UserID)

	if err != nil {
		return "", err
	}

	return user.ID.Hex(), nil
}
