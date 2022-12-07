package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func HandlePostLogin(c *gin.Context) {
	token, err := GenerateToken("Michael")

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
