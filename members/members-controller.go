package members

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleGetMembers(c *gin.Context) {
	members, err := getMembers()

	if err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": members})
}

func handleGetMemberByID(c *gin.Context) {
	member, err := getMemberByID(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": member})
}

// func handlePostMember(c *gin.Context) {
// 	var newMember Member

// 	if err := c.BindJSON(&newMember); err != nil {
// 		return
// 	}

// 	members = append(members, newMember)

// 	c.IndentedJSON(http.StatusCreated, gin.H{"data": newMember.ID})
// }
