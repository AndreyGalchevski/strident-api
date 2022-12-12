package merchandise

import (
	"github.com/AndreyGalchevski/strident-api/auth"
	"github.com/gin-gonic/gin"
)

func InitMerchandiseRouter(router *gin.Engine) {
	router.GET("/merchandise", handleGetMerchandise)
	router.Use(auth.VerifyAuthorization).GET("/merchandise/:id", handleGetMerchandiseByID)
	router.Use(auth.VerifyAuthorization).POST("/merchandise", handlePostMerchandise)
	router.Use(auth.VerifyAuthorization).POST("/merchandise/:id", handlePatchMerchandise)
	router.Use(auth.VerifyAuthorization).DELETE("/merchandise/:id", handleDeleteMerchandise)
}
