package merchandise

import "github.com/gin-gonic/gin"

func InitMerchandiseRouter(router *gin.Engine) {
	router.GET("/merchandise", handleGetMerchandise)
	router.GET("/merchandise/:id", handleGetMerchandiseByID)
	router.POST("/merchandise", handlePostMerchandise)
	router.POST("/merchandise/:id", handlePatchMerchandise)
	router.DELETE("/merchandise/:id", handleDeleteMerchandise)
}
