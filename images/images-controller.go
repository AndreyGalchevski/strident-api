package images

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func HandleUploadImage(c *gin.Context) {
	folderName := c.PostForm("folderName")
	file, _, err := c.Request.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imageURL, err := uploadImage(folderName, file)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": gin.H{"data": imageURL}})
}

type DeleteImageBody struct {
	ImageURL string `json:"imageURL" validate:"required"`
}

func HandleDeleteImage(c *gin.Context) {
	var params DeleteImageBody

	err := c.BindJSON(&params)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validate.Struct(&params)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	err = deleteImage(params.ImageURL)

	if err != nil {
		httpStatus := http.StatusInternalServerError

		if err.Error() == "image not found" {
			httpStatus = http.StatusNotFound
		}

		c.JSON(httpStatus, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"data": gin.H{}})
}
