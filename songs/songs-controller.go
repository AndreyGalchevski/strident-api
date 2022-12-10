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
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": songs})
}

func handleGetSongByID(c *gin.Context) {
	song, err := getSongByID(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": song})
}

func handlePostSong(c *gin.Context) {
	var songData Song

	err := c.BindJSON(&songData)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validate.Struct(&songData)

	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	newSongID, err := createSong(songData)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}

	c.IndentedJSON(http.StatusCreated, gin.H{"data": newSongID})
}
