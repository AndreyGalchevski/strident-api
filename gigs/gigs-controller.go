package gigs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Gig struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Venue   string `json:"venue"`
	Address string `json:"address"`
	City    string `json:"city"`
	Date    string `json:"date"`
	FBEvent string `json:"fbEvent"`
	Image   string `json:"image"`
}

var gigs = []Gig{
	{
		ID:      "1",
		Name:    "Battle For The North",
		Venue:   "Wunderbar",
		Address: "Hativat Golani 18",
		City:    "Haifa",
		Date:    "2022-10-01T17:00:00.000+00:00",
		FBEvent: "https://fb.me/e/2EOm0Nl0z",
		Image:   "https://res.cloudinary.com/dqvimfd8b/image/upload/v1662664964/strident/gigs/production/battle-for-the-north.jpg",
	},
	{
		ID:      "2",
		Name:    "Moshpit Balagan",
		Venue:   "Levontin7",
		Address: "Levontin 7",
		City:    "Tel Aviv",
		Date:    "2020-08-05T18:00:00.000+00:00",
		FBEvent: "https://fb.me/e/1z56ClNSx",
		Image:   "https://res.cloudinary.com/dqvimfd8b/image/upload/v1662664967/strident/gigs/production/moshpit-balagan.jpg",
	},
	{
		ID:      "3",
		Name:    "Rock Show",
		Venue:   "Art Club",
		Address: "Kibutz Galuyot 45",
		City:    "Tel Aviv",
		Date:    "2020-04-14T18:00:00.000+00:00",
		FBEvent: "https://fb.me/e/3bZWPffIY",
		Image:   "https://res.cloudinary.com/dqvimfd8b/image/upload/v1662664964/strident/gigs/production/rock-show.jpg",
	},
	{
		ID:      "4",
		Name:    "Metal 4 Peace",
		Venue:   "Gagarin",
		Address: "Kibutz Galuyot 13",
		City:    "Tel Aviv",
		Date:    "2020-04-02T09:00:00.000+00:00",
		FBEvent: "https://fb.me/e/2sGndwGD9",
		Image:   "https://res.cloudinary.com/dqvimfd8b/image/upload/v1662664966/strident/gigs/production/metal-4-peace.jpg",
	},
	{
		ID:      "5",
		Name:    "Epicenter",
		Venue:   "Jan Jack",
		Address: "Sontsino 17",
		City:    "Tel Aviv",
		Date:    "2020-02-17T18:00:00.000+00:00",
		FBEvent: "https://fb.me/e/1km15UkD6",
		Image:   "https://res.cloudinary.com/dqvimfd8b/image/upload/v1662664969/strident/gigs/production/epi_center.jpg",
	},
	{
		ID:      "6",
		Name:    "YANA ORQO Live Online Fest",
		Venue:   "Peru",
		Address: "World wide",
		City:    "Web",
		Date:    "2020-08-31T18:00:00.000+00:00",
		FBEvent: "",
		Image:   "https://res.cloudinary.com/dqvimfd8b/image/upload/v1596523725/strident/gigs/production/Peru-2020-09-01.jpg",
	},
	{
		ID:      "7",
		Name:    "Bloody News Online Fest",
		Venue:   "Romania",
		Address: "World wide",
		City:    "Web",
		Date:    "2020-08-27T18:00:00.000+00:00",
		FBEvent: "https://www.facebook.com/events/335636320939954/?acontext=%7B%22source%22%3A3%2C%22source_newsfeed_story_type%22%3A%22regular%22%2C%22action_history%22%3A%22%5B%7B%5C%22surface%5C%22%3A%5C%22newsfeed%5C%22%2C%5C%22mechanism%5C%22%3A%5C%22feed_story%5C%22%2C%5C%22extra_data%5C%22%3A%5B%5D%7D%5D%22%2C%22has_source%22%3Atrue%7D&source=3&source_newsfeed_story_type=regular&action_history=%5B%7B%22surface%22%3A%22newsfeed%22%2C%22mechanism%22%3A%22feed_story%22%2C%22extra_data%22%3A%5B%5D%7D%5D&has_source=1&__tn__=K-R&eid=ARBYDQ-JVTL8GNNH3iEZHaMkiv3j9CiLnP1xm4Hd6hThOQqAOPpxdc12BzJODi2FE3zlKmbMcYyhDNq6&fref=mentions&__xts__%5B0%5D=68.ARBR_MVP6pNZ2MBnMk7iLOdqsEL4QAKioaKLdXEBhjsvO9I6z70Bs_n7iKDdejfzuLYPktG9zpHflakZitL6B3ysaD0iwI0LutncBbFkx6-qRyfuqnS5wape-Lex688LZP_tANJw_2a9FEk27zNa4jVnRNBCJ6E4vSeL8uxefXSQGfR9sQDdg6TqHXU9NBMHMTwQUHtehxMMQGAy1ldBP3zfcl-k10kPB72O0ACnCT8XSjLeakqPiVHx0Hj1_tOuxXcmHDZxID6id8mp_IZSUUJvjusY84FCpz_Zgo8yBI4QPYAIllgMn9WCGKB8zBBqj7wWgIIimgsCNGRfjnsf5ZIkWw",
		Image:   "https://res.cloudinary.com/dqvimfd8b/image/upload/v1596523645/strident/gigs/production/Romania-2020-08-28.jpg",
	},
	{
		ID:      "8",
		Name:    "21.2 Nick's B-Day bash & Night of Keoss Strident/03/BullShark",
		Venue:   "Keoss Studios",
		Address: "HaMiktsoa 4",
		City:    "Tel Aviv",
		Date:    "2020-02-21T06:30:00.000+00:00",
		FBEvent: "https://www.facebook.com/events/653666635371023/",
		Image:   "https://res.cloudinary.com/dqvimfd8b/image/upload/v1581252690/strident/gigs/production/KeossStudios-2020-02-21.jpg",
	},
	{
		ID:      "9",
		Name:    "Fuck You Valentine",
		Venue:   "Blaze Rock Bar",
		Address: "Hilel 23",
		City:    "Jerusalem",
		Date:    "2020-02-15T19:30:00.000+00:00",
		FBEvent: "https://www.facebook.com/events/619466475549365/",
		Image:   "https://res.cloudinary.com/dqvimfd8b/image/upload/v1581252534/strident/gigs/production/BlazeRockBar-2020-02-15.jpg",
	},
	{
		ID:      "10",
		Name:    "March Of Plague Release Show",
		Venue:   "Ozen Bar",
		Address: "King George 48",
		City:    "Tel Aviv",
		Date:    "2020-01-16T18:00:00.000+00:00",
		FBEvent: "https://www.facebook.com/events/533651297231767/",
		Image:   "https://res.cloudinary.com/dqvimfd8b/image/upload/v1578320864/strident/gigs/production/OzenBar-%D7%90%D7%95%D7%96%D7%9F%D7%91%D7%A8-2020-01-16.jpg",
	},
	{
		ID:      "11",
		Name:    "Strident ✦ Deusphera ✦ Haifa Double B Studio",
		Venue:   "Double B Studio",
		Address: "Sderot haMeginim 50",
		City:    "Haifa",
		Date:    "2020-01-09T19:00:00.000+00:00",
		FBEvent: "https://www.facebook.com/events/450176279010284/",
		Image:   "https://res.cloudinary.com/dqvimfd8b/image/upload/v1577345561/strident/gigs/production/DoubleBStudio-2009-02-01.jpg",
	},
	{
		ID:      "12",
		Name:    "Metal Market, Winter 2019",
		Venue:   "Ozen Bar",
		Address: "King George 48",
		City:    "Tel Aviv",
		Date:    "2019-12-13T09:00:00.000+00:00",
		FBEvent: "https://facebook.com/events/3078518132222268/?ti=cl",
		Image:   "https://res.cloudinary.com/dqvimfd8b/image/upload/v1572985963/strident/gigs/production/OzenBar-2019-12-13.jpg",
	},
	{
		ID:      "13",
		Name:    "Metal Market, Summer 2019",
		Venue:   "Ozen Bar",
		Address: "King George 48",
		City:    "Tel Aviv",
		Date:    "2019-08-16T08:00:00.000+00:00",
		FBEvent: "https://www.facebook.com/events/359885801359936/",
		Image:   "https://res.cloudinary.com/dqvimfd8b/image/upload/v1571401488/strident/gigs/production/OzenBar-2019-08-16.jpg",
	},
	{
		ID:      "14",
		Name:    "The Show Of Undead, Part 3",
		Venue:   "Art Hall",
		Address: "Ben Avigdor 10",
		City:    "Tel Aviv",
		Date:    "2019-07-12T18:00:00.000+00:00",
		FBEvent: "https://www.facebook.com/events/781053508956086/",
		Image:   "https://res.cloudinary.com/dqvimfd8b/image/upload/v1571401474/strident/gigs/production/ArtHall-2019-07-12.jpg",
	},
	{
		ID:      "15",
		Name:    "The Show Of Undead, Part 2",
		Venue:   "Art Hall",
		Address: "Ben Avigdor 10",
		City:    "Tel Aviv",
		Date:    "2019-06-13T18:00:00.000+00:00",
		FBEvent: "https://www.facebook.com/events/2220730884714744/",
		Image:   "https://res.cloudinary.com/dqvimfd8b/image/upload/v1571401504/strident/gigs/production/ArtHall-2019-06-13.jpg",
	},
	{
		ID:      "16",
		Name:    "The Show Of Undead, part 1",
		Venue:   "Art Hall",
		Address: "Ben Avigdor 10",
		City:    "Tel Aviv",
		Date:    "2019-05-10T17:00:00.000+00:00",
		FBEvent: "https://www.facebook.com/events/491926848006835/?ti=cl",
		Image:   "https://res.cloudinary.com/dqvimfd8b/image/upload/v1572472935/strident/gigs/production/ArtHall-2019-05-10.png",
	},
	{
		ID:      "17",
		Name:    "Death B'Av",
		Venue:   "Gagarin",
		Address: "Shalma 46",
		City:    "Tel Aviv",
		Date:    "2018-07-18T17:00:00.000+00:00",
		FBEvent: "https://www.facebook.com/events/952245078268944/?ti=cl",
		Image:   "https://res.cloudinary.com/dqvimfd8b/image/upload/v1572473665/strident/gigs/production/Gagarin-2018-07-18.png",
	},
}

func handleGetGigs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"data": gigs})
}

func handlePostGig(c *gin.Context) {
	var newGig Gig

	if err := c.BindJSON(&newGig); err != nil {
		return
	}

	gigs = append(gigs, newGig)

	c.IndentedJSON(http.StatusCreated, gin.H{"data": newGig.ID})
}

func handleGetGigByID(c *gin.Context) {
	id := c.Param("id")

	for _, gig := range gigs {
		if gig.ID == id {
			c.IndentedJSON(http.StatusOK, gin.H{"data": gig})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "gig not found"})
}
