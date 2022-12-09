package lyrics

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleGetLyrics(c *gin.Context) {
	lyrics, err := getLyrics()

	if err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": lyrics})
}

func handleGetLyricByID(c *gin.Context) {
	lyric, err := getLyricByID(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": lyric})
}

// func handlePostLyric(c *gin.Context) {
// 	var newLyric Lyric

// 	if err := c.BindJSON(&newLyric); err != nil {
// 		return
// 	}

// 	lyrics = append(lyrics, newLyric)

// 	c.IndentedJSON(http.StatusCreated, gin.H{"data": newLyric.ID})
// }
