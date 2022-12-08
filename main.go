package main

import (
	"os"

	"github.com/AndreyGalchevski/strident-api/auth"
	// Importing the db package in order for the connection to be made
	_ "github.com/AndreyGalchevski/strident-api/db"
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
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("APP_URL")}

	router := gin.Default()

	router.Use(cors.New(config))

	auth.InitAuthRouter(router)
	gigs.InitGigsRouter(router)
	lyrics.InitLyricsRouter(router)
	members.InitMembersRouter(router)
	merchandise.InitMerchandiseRouter(router)
	songs.InitSongsRouter(router)
	videos.InitVideosRouter(router)

	router.Run("localhost:8080")
}
