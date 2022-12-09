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

// func handleGetMerchandiseByID(c *gin.Context) {
// 	id := c.Param("id")

// 	for _, merchandiseItem := range merchandise {
// 		if merchandiseItem.ID == id {
// 			c.IndentedJSON(http.StatusOK, gin.H{"data": merchandiseItem})
// 			return
// 		}
// 	}

// 	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "merchandise not found"})
// }
