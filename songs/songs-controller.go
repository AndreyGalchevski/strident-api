package songs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

// func handlePostSong(c *gin.Context) {
// 	var newSong Song

// 	if err := c.BindJSON(&newSong); err != nil {
// 		return
// 	}

// 	songs = append(songs, newSong)

// 	c.IndentedJSON(http.StatusCreated, gin.H{"data": newSong.ID})
// }
