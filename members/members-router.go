package members

import (
	"github.com/AndreyGalchevski/strident-api/auth"
	"github.com/gin-gonic/gin"
)

func InitMembersRouter(router *gin.Engine) {
	router.GET("/members", handleGetMembers)
	router.Use(auth.VerifyAuthorization).GET("/members/:id", handleGetMemberByID)
	router.Use(auth.VerifyAuthorization).POST("/members", handlePostMember)
	router.Use(auth.VerifyAuthorization).PATCH("/members/:id", handlePatchMember)
	router.Use(auth.VerifyAuthorization).DELETE("/members/:id", handleDeleteMember)
}
