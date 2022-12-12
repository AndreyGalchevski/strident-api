package lyrics

import (
	"github.com/AndreyGalchevski/strident-api/auth"
	"github.com/gin-gonic/gin"
)

func InitLyricsRouter(router *gin.Engine) {
	router.GET("/lyrics", handleGetLyrics)
	router.Use(auth.VerifyAuthorization).GET("/lyrics/:id", handleGetLyricByID)
	router.Use(auth.VerifyAuthorization).POST("/lyrics", handlePostLyric)
	router.Use(auth.VerifyAuthorization).PATCH("/lyrics/:id", handlePatchLyric)
	router.Use(auth.VerifyAuthorization).DELETE("/lyrics/:id", handleDeleteLyric)
}
