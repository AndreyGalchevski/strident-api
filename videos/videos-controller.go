package videos

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Video struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

var videos = []Video{
	{
		ID:   "1",
		Name: "STRIDENT - War (Live At Woko Klub)",
		URL:  "https://www.youtube.com/embed/FKdu9rHsrY0",
	},
	{
		ID:   "2",
		Name: "STRIDENT - No Faith No War (Official video)",
		URL:  "https://www.youtube.com/embed/UkvlRmq62io",
	},
	{
		ID:   "3",
		Name: "Be Metal",
		URL:  "https://www.youtube.com/embed/o0-f8eKQsV8",
	},
	{
		ID:   "4",
		Name: "Nuclear Winter",
		URL:  "https://www.youtube.com/embed/3MpjGJpmG-Y",
	},
	{
		ID:   "5",
		Name: "Final Warhead Blast (Official Lyrics Video)",
		URL:  "https://www.youtube.com/embed/QVKsHgdcT2M",
	},
	{
		ID:   "6",
		Name: "March Of Plague (Official Release)",
		URL:  "https://www.youtube.com/embed/kRaTaPfQi8U",
	},
	{
		ID:   "7",
		Name: "Live In Kiev",
		URL:  "https://www.youtube.com/embed/PWAhdo8AUlI",
	},
	{
		ID:   "8",
		Name: "To Beer Or Not To Beer",
		URL:  "https://youtube.com/embed/tY0U0qxp1Gs",
	},
	{
		ID:   "9",
		Name: "Psycho Provocator",
		URL:  "https://www.youtube.com/embed/599OBkitOYk",
	},
}

func HandleGetVideos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"data": videos})
}

func HandlePostVideo(c *gin.Context) {
	var newVideo Video

	if err := c.BindJSON(&newVideo); err != nil {
		return
	}

	videos = append(videos, newVideo)

	c.IndentedJSON(http.StatusCreated, gin.H{"data": newVideo.ID})
}

func HandleGetVideoByID(c *gin.Context) {
	id := c.Param("id")

	for _, video := range videos {
		if video.ID == id {
			c.IndentedJSON(http.StatusOK, gin.H{"data": video})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "video not found"})
}
