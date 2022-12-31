package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie(AUTH_COOKIE_NAME)

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

		c.Next()
	}
}
