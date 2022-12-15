package members

import (
	// "github.com/AndreyGalchevski/strident-api/auth"
	"github.com/gin-gonic/gin"
)

func InitMembersRouter(r *gin.Engine) {
	r.GET("/members", handleGetMembers)

	authorized := r.Group("/")

	authorized.Use()
	{
		authorized.GET("/members/:id", handleGetMemberByID)
		authorized.POST("/members", handlePostMember)
		authorized.PATCH("/members/:id", handlePatchMember)
		authorized.DELETE("/members/:id", handleDeleteMember)
	}
}
