package members

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Member struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Instrument string `json:"instrument"`
	Image      string `json:"image"`
}

var members = []Member{
	{
		ID:         "1",
		Name:       "Metalych",
		Instrument: "Lead Guitar and Vocals",
		Image:      "https://res.cloudinary.com/dqvimfd8b/image/upload/v1571334257/strident/members/production/misha.jpg",
	},
	{
		ID:         "2",
		Name:       "Yaniv",
		Instrument: "Guitar",
		Image:      "https://res.cloudinary.com/dqvimfd8b/image/upload/v1662666920/strident/members/production/yaniv.jpg",
	},
	{
		ID:         "3",
		Name:       "Artem",
		Instrument: "Bass Guitar",
		Image:      "https://res.cloudinary.com/dqvimfd8b/image/upload/v1571394214/strident/members/production/Artem.jpg",
	},
	{
		ID:         "4",
		Name:       "Andrey",
		Instrument: "Drums",
		Image:      "https://res.cloudinary.com/dqvimfd8b/image/upload/v1571394243/strident/members/production/Andrey.jpg",
	},
}

func HandleGetMembers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"data": members})
}

func HandlePostMember(c *gin.Context) {
	var newMember Member

	if err := c.BindJSON(&newMember); err != nil {
		return
	}

	members = append(members, newMember)

	c.IndentedJSON(http.StatusCreated, gin.H{"data": newMember.ID})
}

func HandleGetMemberByID(c *gin.Context) {
	id := c.Param("id")

	for _, member := range members {
		if member.ID == id {
			c.IndentedJSON(http.StatusOK, gin.H{"data": member})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "member not found"})
}
