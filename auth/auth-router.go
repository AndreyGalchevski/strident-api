package auth

import "github.com/gin-gonic/gin"

func InitAuthRouter(router *gin.Engine) {
	router.POST("/auth/login", handlePostLogin)
	router.GET("/auth/verify", handleGetVerify)
}
