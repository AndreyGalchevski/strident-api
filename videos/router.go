package videos

import (
	"github.com/AndreyGalchevski/strident-api/auth"
	"github.com/gorilla/mux"
)

func InitVideosRouter(r *mux.Router) {
	r.HandleFunc("/videos", handleGetVideos).Methods("GET")

	r.HandleFunc("/videos/{id}", auth.VerifyAuthorization(handleGetVideoByID)).Methods("GET")
	r.HandleFunc("/videos", auth.VerifyAuthorization(handlePostVideo)).Methods("POST")
	r.HandleFunc("/videos/{id}", auth.VerifyAuthorization(handlePatchVideo)).Methods("PATCH")
	r.HandleFunc("/videos/{id}", auth.VerifyAuthorization(handleDeleteVideo)).Methods("DELETE")
}
