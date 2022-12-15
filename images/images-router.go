package images

import (
	"github.com/AndreyGalchevski/strident-api/auth"
	"github.com/gin-gonic/gin"
)

func InitImagesRouter(r *gin.Engine) {
	authorized := r.Group("/")

	authorized.Use(auth.VerifyAuthorization())
	{
		authorized.POST("/images", handlePostImage)
		authorized.DELETE("/images", handleDeleteImage)
	}
}
