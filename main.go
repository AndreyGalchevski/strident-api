package main

import (
	"net/http"
	"os"

	"github.com/AndreyGalchevski/strident-api/auth"
	"github.com/AndreyGalchevski/strident-api/db"
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

func handlePreflight() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, gin.H{})
			return
		}

		c.Next()
	}
}

func main() {
	err := godotenv.Load()

	if err != nil && err.Error() != "open .env: no such file or directory" {
		panic("Error loading .env file: " + err.Error())
	}

	db.Connect()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("WEB_APP_URL")}
	config.AllowCredentials = true

	if os.Getenv("APP_ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.SetTrustedProxies(nil)

	router.MaxMultipartMemory = 8 << 20

	router.Use(cors.New(config))

	router.Use(handlePreflight())

	auth.InitAuthRouter(router)
	gigs.InitGigsRouter(router)
	lyrics.InitLyricsRouter(router)
	members.InitMembersRouter(router)
	merchandise.InitMerchandiseRouter(router)
	songs.InitSongsRouter(router)
	videos.InitVideosRouter(router)

	router.Run()
}
