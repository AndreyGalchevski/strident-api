package videos

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleGetVideos(c *gin.Context) {
	videos, err := getVideos()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": videos})
}

func handleGetVideoByID(c *gin.Context) {
	video, err := getVideoByID(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": video})
}

// func handlePostVideo(c *gin.Context) {
// 	var newVideo Video

// 	if err := c.BindJSON(&newVideo); err != nil {
// 		return
// 	}

// 	videos = append(videos, newVideo)

// 	c.IndentedJSON(http.StatusCreated, gin.H{"data": newVideo.ID})
// }
