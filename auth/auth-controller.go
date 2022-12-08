package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type User struct {
	username string
	password string
}

var user User = User{username: "michael123", password: "123"}

func HandlePostLogin(c *gin.Context) {
	var credentials Credentials

	if err := c.BindJSON(&credentials); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Please try again"})
		return
	}

	// TODO: fetch the user from the db

	if credentials.Password != user.password {
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

func HandlePostVerify(c *gin.Context) {
	username, err := VerifyToken(c.Param("token"))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": username})
}
