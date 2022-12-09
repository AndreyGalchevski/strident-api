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

// func handlePostSong(c *gin.Context) {
// 	var newSong Song

// 	if err := c.BindJSON(&newSong); err != nil {
// 		return
// 	}

// 	songs = append(songs, newSong)

// 	c.IndentedJSON(http.StatusCreated, gin.H{"data": newSong.ID})
// }

// func handleGetSongByID(c *gin.Context) {
// 	id := c.Param("id")

// 	for _, song := range songs {
// 		if song.ID == id {
// 			c.IndentedJSON(http.StatusOK, gin.H{"data": song})
// 			return
// 		}
// 	}

// 	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "song not found"})
// }
