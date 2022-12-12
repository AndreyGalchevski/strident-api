package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyAuthorization(c *gin.Context) {
	cookie, err := c.Cookie("stridentToken")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	_, err = VerifyToken(cookie)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
}
