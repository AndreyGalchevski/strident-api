package lyrics

import (
	"github.com/AndreyGalchevski/strident-api/auth"
	"github.com/gorilla/mux"
)

func InitLyricsRouter(r *mux.Router) {
	r.HandleFunc("/lyrics", handleGetLyrics).Methods("GET")

	r.HandleFunc("/lyrics/{id}", auth.VerifyAuthorization(handleGetLyricByID)).Methods("GET")
	r.HandleFunc("/lyrics", auth.VerifyAuthorization(handlePostLyric)).Methods("POST")
	r.HandleFunc("/lyrics/{id}", auth.VerifyAuthorization(handlePatchLyric)).Methods("PATCH")
	r.HandleFunc("/lyrics/{id}", auth.VerifyAuthorization(handleDeleteLyric)).Methods("DELETE")
}
