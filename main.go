package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/AndreyGalchevski/strident-api/auth"
	"github.com/AndreyGalchevski/strident-api/db"
	"github.com/AndreyGalchevski/strident-api/gigs"
	"github.com/AndreyGalchevski/strident-api/lyrics"
	"github.com/AndreyGalchevski/strident-api/members"
	"github.com/AndreyGalchevski/strident-api/merchandise"
	"github.com/AndreyGalchevski/strident-api/songs"
	"github.com/AndreyGalchevski/strident-api/videos"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	err := godotenv.Load()

	if err != nil && err.Error() != "open .env: no such file or directory" {
		panic("Error loading .env file: " + err.Error())
	}

	db.Connect()

	if os.Getenv("APP_ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.SetTrustedProxies(nil)

	router.MaxMultipartMemory = 8 << 20

	corsConfig := cors.Options{
		AllowedOrigins:     []string{strings.TrimSuffix(os.Getenv("WEB_APP_URL"), "/")},
		AllowCredentials:   true,
		OptionsPassthrough: false,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		Debug: true,
	}

	router.Use(cors.New(corsConfig))

	auth.InitAuthRouter(router)
	gigs.InitGigsRouter(router)
	lyrics.InitLyricsRouter(router)
	members.InitMembersRouter(router)
	merchandise.InitMerchandiseRouter(router)
	songs.InitSongsRouter(router)
	videos.InitVideosRouter(router)

	router.Run()
}
