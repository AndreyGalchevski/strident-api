package gigs

import (
	"github.com/AndreyGalchevski/strident-api/auth"
	"github.com/gin-gonic/gin"
)

func InitGigsRouter(r *gin.Engine) {
	r.GET("/gigs", handleGetGigs)

	authorized := r.Group("/")

	authorized.Use(auth.VerifyAuthorization())
	{
		authorized.GET("/gigs/:id", handleGetGigByID)
		authorized.POST("/gigs", handlePostGig)
		authorized.PATCH("/gigs/:id", handlePatchGig)
		authorized.DELETE("/gigs/:id", handleDeleteGig)
	}
}
