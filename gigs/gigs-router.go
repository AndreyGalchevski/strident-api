package gigs

import "github.com/gin-gonic/gin"

func InitGigsRouter(router *gin.Engine) {
	router.GET("/gigs", handleGetGigs)
	router.GET("/gigs/:id", handleGetGigByID)
	router.POST("/gigs", handlePostGig)
}
