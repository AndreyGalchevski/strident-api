package videos

import (
	"github.com/AndreyGalchevski/strident-api/auth"
	"github.com/gin-gonic/gin"
)

func InitVideosRouter(r *gin.Engine) {
	r.GET("/videos", handleGetVideos)

	authorized := r.Group("/")

	authorized.Use(auth.VerifyAuthorization())
	{
		authorized.GET("/videos/:id", handleGetVideoByID)
		authorized.POST("/videos", handlePostVideo)
		authorized.PATCH("/videos/:id", handlePatchVideo)
		authorized.DELETE("/videos/:id", handleDeleteVideo)
	}
}
