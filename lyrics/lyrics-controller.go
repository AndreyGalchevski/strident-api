package lyrics

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleGetLyrics(c *gin.Context) {
	lyrics, err := getLyrics()

	if err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": lyrics})
}

// func handlePostLyric(c *gin.Context) {
// 	var newLyric Lyric

// 	if err := c.BindJSON(&newLyric); err != nil {
// 		return
// 	}

// 	lyrics = append(lyrics, newLyric)

// 	c.IndentedJSON(http.StatusCreated, gin.H{"data": newLyric.ID})
// }

// func handleGetLyricByID(c *gin.Context) {
// 	id := c.Param("id")

// 	for _, lyric := range lyrics {
// 		if lyric.ID == id {
// 			c.IndentedJSON(http.StatusOK, gin.H{"data": lyric})
// 			return
// 		}
// 	}

// 	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "lyric not found"})
// }
