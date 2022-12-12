package videos

import (
	"github.com/AndreyGalchevski/strident-api/auth"
	"github.com/gin-gonic/gin"
)

func InitVideosRouter(router *gin.Engine) {
	router.GET("/videos", handleGetVideos)
	router.Use(auth.VerifyAuthorization).GET("/videos/:id", handleGetVideoByID)
	router.Use(auth.VerifyAuthorization).POST("/videos", handlePostVideo)
	router.Use(auth.VerifyAuthorization).PATCH("/videos/:id", handlePatchVideo)
	router.Use(auth.VerifyAuthorization).DELETE("/videos/:id", handleDeleteVideo)
}
