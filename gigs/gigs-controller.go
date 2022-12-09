package gigs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

// func handlePostGig(c *gin.Context) {
// 	var newGig Gig

// 	if err := c.BindJSON(&newGig); err != nil {
// 		return
// 	}

// 	gigs = append(gigs, newGig)

// 	c.IndentedJSON(http.StatusCreated, gin.H{"data": newGig.ID})
// }
