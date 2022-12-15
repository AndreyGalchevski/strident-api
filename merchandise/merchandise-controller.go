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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": merchandise})
}

func handleGetMerchandiseByID(c *gin.Context) {
	merchandise, err := getMerchandiseByID(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": merchandise})
}

func handlePostMerchandise(c *gin.Context) {
	var params MerchandiseFormData

	err := c.Bind(&params)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validate.Struct(&params)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	image, _, err := c.Request.FormFile("image")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newMerchandiseID, err := createMerchandise(params, image)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}

	c.JSON(http.StatusCreated, gin.H{"data": newMerchandiseID})
}

func handlePatchMerchandise(c *gin.Context) {
	var params MerchandiseFormData

	err := c.Bind(&params)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validate.Struct(&params)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ok, err := updateMerchandise(c.Param("id"), params)

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

func handleDeleteMerchandise(c *gin.Context) {
	ok, err := deleteMerchandise(c.Param("id"))

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
