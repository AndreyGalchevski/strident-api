package auth

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

const AUTH_COOKIE_NAME = "stridentToken"

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func handlePostLogin(c *gin.Context) {
	var credentials Credentials

	err := c.BindJSON(&credentials)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please try again"})
		return
	}

	token, err := login(credentials)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	isProd := os.Getenv("APP_ENV") == "prod"

	if isProd {
		c.SetSameSite(http.SameSiteNoneMode)
	}

	domain := ""

	if isProd {
		domain = strings.TrimSuffix(os.Getenv("WEB_APP_URL"), "/")
	}

	c.SetCookie(
		AUTH_COOKIE_NAME,
		token,
		int(TokenMaxAge.Seconds()),
		"/",
		domain,
		isProd,
		true,
	)
}

func handleGetVerify(c *gin.Context) {
	_, err := c.Cookie(AUTH_COOKIE_NAME)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Session expired"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"data": gin.H{}})
}
