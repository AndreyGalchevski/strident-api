package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func handlePostLogin(c *gin.Context) {
	var credentials Credentials

	if err := c.BindJSON(&credentials); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Please try again"})
		return
	}

	token, err := login(credentials)

	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
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
