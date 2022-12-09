package videos

import "github.com/gin-gonic/gin"

func InitVideosRouter(router *gin.Engine) {
	router.GET("/videos", handleGetVideos)
	router.GET("/videos/:id", handleGetVideoByID)
	// router.POST("/videos", handlePostVideo)
}
