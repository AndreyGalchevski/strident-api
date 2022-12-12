package gigs

import (
	"github.com/AndreyGalchevski/strident-api/auth"
	"github.com/gin-gonic/gin"
)

func InitGigsRouter(router *gin.Engine) {
	router.GET("/gigs", handleGetGigs)
	router.Use(auth.VerifyAuthorization).GET("/gigs/:id", handleGetGigByID)
	router.Use(auth.VerifyAuthorization).POST("/gigs", handlePostGig)
	router.Use(auth.VerifyAuthorization).PATCH("/gigs/:id", handlePatchGig)
	router.Use(auth.VerifyAuthorization).DELETE("/gigs/:id", handleDeleteGig)
}
