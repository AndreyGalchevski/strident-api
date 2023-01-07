package songs

import (
	"github.com/AndreyGalchevski/strident-api/auth"
	"github.com/gorilla/mux"
)

func InitSongsRouter(r *mux.Router) {
	r.HandleFunc("/songs", handleGetSongs).Methods("GET")

	r.HandleFunc("/songs/{id}", auth.VerifyAuthorization(handleGetSongByID)).Methods("GET")
	r.HandleFunc("/songs", auth.VerifyAuthorization(handlePostSong)).Methods("POST")
	r.HandleFunc("/songs/{id}", auth.VerifyAuthorization(handlePatchSong)).Methods("PATCH")
	r.HandleFunc("/songs/{id}", auth.VerifyAuthorization(handleDeleteSong)).Methods("DELETE")
}
