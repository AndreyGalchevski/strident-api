package songs

import (
	"github.com/AndreyGalchevski/strident-api/auth"
	"github.com/gin-gonic/gin"
)

func InitSongsRouter(router *gin.Engine) {
	router.GET("/songs", handleGetSongs)
	router.Use(auth.VerifyAuthorization).GET("/songs/:id", handleGetSongByID)
	router.Use(auth.VerifyAuthorization).POST("/songs", handlePostSong)
	router.Use(auth.VerifyAuthorization).PATCH("/songs/:id", handlePatchSong)
	router.Use(auth.VerifyAuthorization).DELETE("/songs/:id", handleDeleteSong)
}
