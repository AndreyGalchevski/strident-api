package merchandise

import (
	"github.com/AndreyGalchevski/strident-api/auth"
	"github.com/gin-gonic/gin"
)

func InitMerchandiseRouter(r *gin.Engine) {
	r.GET("/merchandise", handleGetMerchandise)

	authorized := r.Group("/")

	authorized.Use(auth.VerifyAuthorization())
	{
		authorized.GET("/merchandise/:id", handleGetMerchandiseByID)
		authorized.POST("/merchandise", handlePostMerchandise)
		authorized.PATCH("/merchandise/:id", handlePatchMerchandise)
		authorized.DELETE("/merchandise/:id", handleDeleteMerchandise)
	}
}
