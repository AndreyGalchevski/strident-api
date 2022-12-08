package songs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Song struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	URL   string `json:"url"`
	Album string `json:"album"`
}

var songs = []Song{
	{
		ID:    "1",
		Name:  "No Faith No War",
		URL:   "https://open.spotify.com/embed/track/3E4zDBJCas8gy5YINUchFz",
		Album: "March Of Plague",
	},
	{
		ID:    "2",
		Name:  "March Of Plague",
		URL:   "https://open.spotify.com/embed/track/0OaTAp8LVe7doqiB7sggCs",
		Album: "March Of Plague",
	},
	{
		ID:    "3",
		Name:  "Be Metal",
		URL:   "https://open.spotify.com/embed/track/1jElrQH7U55eckVkOrnfSI",
		Album: "March Of Plague",
	},
	{
		ID:    "4",
		URL:   "https://open.spotify.com/embed/track/4NRqH03A4Rb0mSSejSb56t",
		Name:  "Psycho Provocator",
		Album: "On The Aim",
	},
	{
		ID:    "5",
		Name:  "Prepare To Die",
		URL:   "https://open.spotify.com/embed/track/2fgu6pVLPBUtsj7z8sJJq0",
		Album: "On The Aim",
	},
	{
		ID:    "6",
		Name:  "To Beer Or Not To Beer",
		URL:   "https://open.spotify.com/embed/track/1Njqc3kRDnwBFZiMXeEoJS",
		Album: "On The Aim",
	},
}

func handleGetSongs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"data": songs})
}

func handlePostSong(c *gin.Context) {
	var newSong Song

	if err := c.BindJSON(&newSong); err != nil {
		return
	}

	songs = append(songs, newSong)

	c.IndentedJSON(http.StatusCreated, gin.H{"data": newSong.ID})
}

func handleGetSongByID(c *gin.Context) {
	id := c.Param("id")

	for _, song := range songs {
		if song.ID == id {
			c.IndentedJSON(http.StatusOK, gin.H{"data": song})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "song not found"})
}
