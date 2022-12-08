package songs

import "github.com/gin-gonic/gin"

func InitSongsRouter(router *gin.Engine) {
	router.GET("/songs", handleGetSongs)
	router.GET("/songs/:id", handleGetSongByID)
	router.POST("/songs", handlePostSong)
}
