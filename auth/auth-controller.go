package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/AndreyGalchevski/strident-api/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

var usersCollection *mongo.Collection = db.GetCollection(db.DBClient, "users")

func handlePostLogin(c *gin.Context) {
	var credentials Credentials

	if err := c.BindJSON(&credentials); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Please try again"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User

	err := usersCollection.FindOne(ctx, bson.M{"username": credentials.Username}).Decode(&user)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Wrong credentials"})
			return
		}

		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if credentials.Password != user.Password {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Wrong credentials"})
		return
	}

	token, err := GenerateToken(credentials.Username)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": token})
}

func handlePostVerify(c *gin.Context) {
	username, err := VerifyToken(c.Param("token"))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": username})
}
