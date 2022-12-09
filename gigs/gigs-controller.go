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

// func handlePostGig(c *gin.Context) {
// 	var newGig Gig

// 	if err := c.BindJSON(&newGig); err != nil {
// 		return
// 	}

// 	gigs = append(gigs, newGig)

// 	c.IndentedJSON(http.StatusCreated, gin.H{"data": newGig.ID})
// }

// func handleGetGigByID(c *gin.Context) {
// 	id := c.Param("id")

// 	for _, gig := range gigs {
// 		if gig.ID == id {
// 			c.IndentedJSON(http.StatusOK, gin.H{"data": gig})
// 			return
// 		}
// 	}

// 	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "gig not found"})
// }
