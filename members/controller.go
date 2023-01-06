package members

import (
	"net/http"

	"github.com/AndreyGalchevski/strident-api/db"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func handleGetMembers(c *gin.Context) {
	members, err := getMembers()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": members})
}

func handleGetMemberByID(c *gin.Context) {
	member, err := getMemberByID(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": member})
}

func handlePostMember(c *gin.Context) {
	var params db.Member

	err := c.Bind(&params)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validate.Struct(&params)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Please fill out all the required fields"})
		return
	}

	image, _, err := c.Request.FormFile("image")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newMemberID, err := createMember(params, image)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}

	c.JSON(http.StatusCreated, gin.H{"data": newMemberID})
}

func handlePatchMember(c *gin.Context) {
	var params db.Member

	err := c.Bind(&params)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validate.Struct(&params)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Please fill out all the required fields"})
		return
	}

	image, _, _ := c.Request.FormFile("image")

	ok, err := updateMember(c.Param("id"), params, image)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"data": gin.H{}})
}

func handleDeleteMember(c *gin.Context) {
	ok, err := deleteMember(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"data": gin.H{}})
}
