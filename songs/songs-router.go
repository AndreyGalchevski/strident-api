package songs

import (
	"github.com/AndreyGalchevski/strident-api/auth"
	"github.com/gin-gonic/gin"
)

func InitSongsRouter(r *gin.Engine) {
	r.GET("/songs", handleGetSongs)

	authorized := r.Group("/")

	authorized.Use(auth.VerifyAuthorization())
	{
		authorized.GET("/songs/:id", handleGetSongByID)
		authorized.POST("/songs", handlePostSong)
		authorized.PATCH("/songs/:id", handlePatchSong)
		authorized.DELETE("/songs/:id", handleDeleteSong)
	}
}
