package members

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

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

func handlePostMember(c *gin.Context) {
	var memberData Member

	err := c.BindJSON(&memberData)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validate.Struct(&memberData)

	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	newMemberID, err := createMember(memberData)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}

	c.IndentedJSON(http.StatusCreated, gin.H{"data": newMemberID})
}
