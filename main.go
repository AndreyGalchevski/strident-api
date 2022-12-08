package main

import (
	"os"

	"github.com/AndreyGalchevski/strident-api/auth"
	"github.com/AndreyGalchevski/strident-api/gigs"
	"github.com/AndreyGalchevski/strident-api/lyrics"
	"github.com/AndreyGalchevski/strident-api/members"
	"github.com/AndreyGalchevski/strident-api/merchandise"
	"github.com/AndreyGalchevski/strident-api/songs"
	"github.com/AndreyGalchevski/strident-api/videos"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("APP_URL")}

	router.Use(cors.New(config))

	router.POST("/auth/login", auth.HandlePostLogin)

	router.GET("/gigs", gigs.HandleGetGigs)
	router.GET("/gigs/:id", gigs.HandleGetGigByID)
	router.POST("/gigs", gigs.HandlePostGig)

	router.GET("/lyrics", lyrics.HandleGetLyrics)
	router.GET("/lyrics/:id", lyrics.HandleGetLyricByID)
	router.POST("/lyrics", lyrics.HandlePostLyric)

	router.GET("/members", members.HandleGetMembers)
	router.GET("/members/:id", members.HandleGetMemberByID)
	router.POST("/members", members.HandlePostMember)

	router.GET("/merchandise", merchandise.HandleGetMerchandise)
	router.GET("/merchandise/:id", merchandise.HandleGetMerchandiseByID)
	router.POST("/merchandise", merchandise.HandlePostMerchandise)

	router.GET("/songs", songs.HandleGetSongs)
	router.GET("/songs/:id", songs.HandleGetSongByID)
	router.POST("/songs", songs.HandlePostSong)

	router.GET("/videos", videos.HandleGetVideos)
	router.GET("/videos/:id", videos.HandleGetVideoByID)
	router.POST("/videos", videos.HandlePostVideo)

	router.POST("/login", auth.HandlePostLogin)
	router.POST("/verify/:token", auth.HandlePostVerify)

	router.Run("localhost:8080")
}
