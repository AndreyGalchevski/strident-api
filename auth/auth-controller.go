package auth

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func handlePostLogin(c *gin.Context) {
	var credentials Credentials

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please try again"})
		return
	}

	token, err := login(credentials)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if os.Getenv("APP_ENV") == "production" {
		c.SetSameSite(http.SameSiteStrictMode)
	}

	c.SetCookie(
		"stridentToken",
		token,
		int(TokenMaxAge.Seconds()),
		"",
		"",
		os.Getenv("APP_ENV") == "production",
		true,
	)
}

func handlePostVerify(c *gin.Context) {
	username, err := VerifyToken(c.Param("token"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": username})
}
