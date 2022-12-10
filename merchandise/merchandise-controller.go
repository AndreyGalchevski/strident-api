package merchandise

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func handleGetMerchandise(c *gin.Context) {
	merchandise, err := getMerchandise()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": merchandise})
}

func handleGetMerchandiseByID(c *gin.Context) {
	merchandise, err := getMerchandiseByID(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": merchandise})
}

func handlePostMerchandise(c *gin.Context) {
	var merchandiseData Merchandise

	err := c.BindJSON(&merchandiseData)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validate.Struct(&merchandiseData)

	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	newMerchandiseID, err := createMerchandise(merchandiseData)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}

	c.IndentedJSON(http.StatusCreated, gin.H{"data": newMerchandiseID})
}
