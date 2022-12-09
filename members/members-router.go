package members

import "github.com/gin-gonic/gin"

func InitMembersRouter(router *gin.Engine) {
	router.GET("/members", handleGetMembers)
	router.GET("/members/:id", handleGetMemberByID)
	// router.POST("/members", handlePostMember)
}
