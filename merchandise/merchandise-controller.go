package merchandise

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Merchandise struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"` // TODO: use enum "Digital album" | "CD" | "T-shirt" | "Girls T-shirt" | "Patch"
	Price int    `json:"price"`
	URL   string `json:"url"`
	Image string `json:"image"`
}

var merchandise = []Merchandise{
	{
		ID:    "1",
		Name:  "March Of Plague",
		Type:  "Digital album",
		Price: 700,
		URL:   "https://stridentthrash.bandcamp.com/album/march-of-plague",
		Image: "https://res.cloudinary.com/dqvimfd8b/image/upload/v1574771723/strident/merchandises/production/March%20Of%20Plague-Digital%20Album.jpg",
	},
	{
		ID:    "2",
		Name:  "March Of Plague",
		Type:  "CD",
		Price: 1600,
		URL:   "https://stridentthrash.bandcamp.com/album/march-of-plague",
		Image: "https://res.cloudinary.com/dqvimfd8b/image/upload/v1577700237/strident/merchandises/production/March%20Of%20Plague-CD.jpg",
	},
	{
		ID:    "3",
		Name:  "On The Aim",
		Type:  "Digital album",
		Price: 400,
		URL:   "https://stridentthrash.bandcamp.com/album/on-the-aim",
		Image: "https://res.cloudinary.com/dqvimfd8b/image/upload/v1577700537/strident/merchandises/production/On%20The%20Aim-Digital%20album.jpg",
	},
	{
		ID:    "4",
		Name:  "On The Aim",
		Type:  "CD",
		Price: 1350,
		URL:   "https://stridentthrash.bandcamp.com/album/on-the-aim",
		Image: "https://res.cloudinary.com/dqvimfd8b/image/upload/v1577700588/strident/merchandises/production/On%20The%20Aim-CD.jpg",
	},
	{
		ID:    "5",
		Name:  "March Of Plague (Two sided)",
		Type:  "T-shirt",
		Price: 2000,
		URL:   "https://stridentthrash.bandcamp.com/merch/mans-two-sided-print-t-shirt-weeder-new-design",
		Image: "https://res.cloudinary.com/dqvimfd8b/image/upload/v1577700792/strident/merchandises/production/March%20Of%20Plague%20man%27s%20two%20sided%20print-t-shirt.jpg",
	},
	{
		ID:    "6",
		Name:  "March Of Plague (Two sided)",
		Type:  "Girls T-shirt",
		Price: 2000,
		URL:   "https://stridentthrash.bandcamp.com/merch/womans-two-sided-print-t-shirt-weeder-new-design",
		Image: "https://res.cloudinary.com/dqvimfd8b/image/upload/v1577701565/strident/merchandises/production/March%20Of%20Plague%20woman%27s%20two%20sided%20print-t-shirt.jpg",
	},
	{
		ID:    "7",
		Name:  "March Of Plague (One sided)",
		Type:  "T-shirt",
		Price: 1700,
		URL:   "https://stridentthrash.bandcamp.com/merch/mans-one-side-print-t-shirt-march-of-plague-limited-edition",
		Image: "https://res.cloudinary.com/dqvimfd8b/image/upload/v1577700917/strident/merchandises/production/Limited%20edition%20March%20Of%20Plague%20man%27s%20one%20side%20print-t-shirt.jpg",
	},
	{
		ID:    "8",
		Name:  "March Of Plague (One sided)",
		Type:  "Girls T-shirt",
		Price: 1700,
		URL:   "https://stridentthrash.bandcamp.com/merch/womans-one-side-print-t-shirt-march-of-plague-limited-edition",
		Image: "https://res.cloudinary.com/dqvimfd8b/image/upload/v1577701798/strident/merchandises/production/Limited%20edition%20March%20Of%20Plague%20Woman%27s%20one%20side%20print-t-shirt.jpg",
	},
	{
		ID:    "9",
		Name:  "Embroidered logo",
		Type:  "Patch",
		Price: 500,
		URL:   "https://stridentthrash.bandcamp.com/merch/embroidered-logo-patch",
		Image: "https://res.cloudinary.com/dqvimfd8b/image/upload/v1577701975/strident/merchandises/production/Embroidered%20logo-patch.jpg",
	},
}

func HandleGetMerchandise(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"data": merchandise})
}

func HandlePostMerchandise(c *gin.Context) {
	var newMerchandise Merchandise

	if err := c.BindJSON(&newMerchandise); err != nil {
		return
	}

	merchandise = append(merchandise, newMerchandise)

	c.IndentedJSON(http.StatusCreated, gin.H{"data": newMerchandise.ID})
}

func HandleGetMerchandiseByID(c *gin.Context) {
	id := c.Param("id")

	for _, merchandiseItem := range merchandise {
		if merchandiseItem.ID == id {
			c.IndentedJSON(http.StatusOK, gin.H{"data": merchandiseItem})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "merchandise not found"})
}
