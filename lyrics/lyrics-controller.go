package lyrics

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func handleGetLyrics(c *gin.Context) {
	lyrics, err := getLyrics()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": lyrics})
}

func handleGetLyricByID(c *gin.Context) {
	lyric, err := getLyricByID(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": lyric})
}

func handlePostLyric(c *gin.Context) {
	var params LyricFormData

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

	newLyricID, err := createLyric(params)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}

	c.JSON(http.StatusCreated, gin.H{"data": newLyricID})
}

func handlePatchLyric(c *gin.Context) {
	var lyricData Lyric

	err := c.Bind(&lyricData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validate.Struct(&lyricData)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ok, err := updateLyric(c.Param("id"), lyricData)

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

func handleDeleteLyric(c *gin.Context) {
	ok, err := deleteLyric(c.Param("id"))

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
