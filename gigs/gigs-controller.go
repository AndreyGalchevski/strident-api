package gigs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func handleGetGigs(c *gin.Context) {
	gigs, err := getGigs()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": gigs})
}

func handleGetGigByID(c *gin.Context) {
	gig, err := getGigByID(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": gig})
}

func handlePostGig(c *gin.Context) {
	var gigData Gig

	err := c.BindJSON(&gigData)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validate.Struct(&gigData)

	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	newGigID, err := createGig(gigData)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}

	c.IndentedJSON(http.StatusCreated, gin.H{"data": newGigID})
}
