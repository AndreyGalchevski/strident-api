package merchandise

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleGetMerchandise(c *gin.Context) {
	merchandise, err := getMerchandise()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": merchandise})
}

// func handlePostMerchandise(c *gin.Context) {
// 	var newMerchandise Merchandise

// 	if err := c.BindJSON(&newMerchandise); err != nil {
// 		return
// 	}

// 	merchandise = append(merchandise, newMerchandise)

// 	c.IndentedJSON(http.StatusCreated, gin.H{"data": newMerchandise.ID})
// }

func handleGetMerchandiseByID(c *gin.Context) {
	merchandise, err := getMerchandiseByID(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": merchandise})
}
