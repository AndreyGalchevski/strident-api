package lyrics

import "github.com/gin-gonic/gin"

func InitLyricsRouter(router *gin.Engine) {
	router.GET("/lyrics", handleGetLyrics)
	router.GET("/lyrics/:id", handleGetLyricByID)
	router.POST("/lyrics", handlePostLyric)
	router.PATCH("/lyrics/:id", handlePatchLyric)
}
