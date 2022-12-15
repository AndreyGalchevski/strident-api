package songs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func handleGetSongs(c *gin.Context) {
	songs, err := getSongs()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": songs})
}

func handleGetSongByID(c *gin.Context) {
	song, err := getSongByID(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": song})
}

func handlePostSong(c *gin.Context) {
	var params SongFormData

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

	newSongID, err := createSong(params)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}

	c.JSON(http.StatusCreated, gin.H{"data": newSongID})
}

func handlePatchSong(c *gin.Context) {
	var params SongFormData

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

	ok, err := updateSong(c.Param("id"), params)

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

func handleDeleteSong(c *gin.Context) {
	ok, err := deleteSong(c.Param("id"))

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
