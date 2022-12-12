package lyrics

import (
	"github.com/AndreyGalchevski/strident-api/auth"
	"github.com/gin-gonic/gin"
)

func InitLyricsRouter(r *gin.Engine) {
	r.GET("/lyrics", handleGetLyrics)

	authorized := r.Group("/")

	authorized.Use(auth.VerifyAuthorization())
	{
		authorized.GET("/lyrics/:id", handleGetLyricByID)
		authorized.POST("/lyrics", handlePostLyric)
		authorized.PATCH("/lyrics/:id", handlePatchLyric)
		authorized.DELETE("/lyrics/:id", handleDeleteLyric)
	}
}
