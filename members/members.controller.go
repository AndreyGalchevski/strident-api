package members

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleGetMembers(c *gin.Context) {
	members, err := getMembers()

	if err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": members})
}

// func handlePostMember(c *gin.Context) {
// 	var newMember Member

// 	if err := c.BindJSON(&newMember); err != nil {
// 		return
// 	}

// 	members = append(members, newMember)

// 	c.IndentedJSON(http.StatusCreated, gin.H{"data": newMember.ID})
// }

// func handleGetMemberByID(c *gin.Context) {
// 	id := c.Param("id")

// 	for _, member := range members {
// 		if member.ID == id {
// 			c.IndentedJSON(http.StatusOK, gin.H{"data": member})
// 			return
// 		}
// 	}

// 	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "member not found"})
// }
