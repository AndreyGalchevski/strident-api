package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/AndreyGalchevski/strident-api/auth"
	"github.com/AndreyGalchevski/strident-api/db"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/AndreyGalchevski/strident-api/albums"
	"github.com/AndreyGalchevski/strident-api/gigs"
	"github.com/AndreyGalchevski/strident-api/lyrics"
	"github.com/AndreyGalchevski/strident-api/members"
	"github.com/AndreyGalchevski/strident-api/merchandise"
	"github.com/AndreyGalchevski/strident-api/songs"
	"github.com/AndreyGalchevski/strident-api/videos"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil && err.Error() != "open .env: no such file or directory" {
		panic("Error loading .env file: " + err.Error())
	}

	db.Connect()

	corsConfig := cors.Options{
		AllowedOrigins:     []string{os.Getenv("WEB_APP_URL")},
		AllowCredentials:   true,
		OptionsPassthrough: false,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
	}

	router := mux.NewRouter()

	auth.InitAuthRouter(router)
	albums.InitAlbumsRouter(router)
	gigs.InitGigsRouter(router)
	lyrics.InitLyricsRouter(router)
	members.InitMembersRouter(router)
	merchandise.InitMerchandiseRouter(router)
	songs.InitSongsRouter(router)
	videos.InitVideosRouter(router)

	routerWithCors := cors.New(corsConfig).Handler(router)

	http.ListenAndServe(":8080", routerWithCors)

	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server closed\n")
	} else if err != nil {
		log.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
