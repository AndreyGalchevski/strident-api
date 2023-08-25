package albums

import (
	"github.com/AndreyGalchevski/strident-api/auth"
	"github.com/gorilla/mux"
)

func InitAlbumsRouter(r *mux.Router) {
	r.HandleFunc("/albums", handleGetAlbums).Methods("GET")

	r.HandleFunc("/albums/{id}", auth.VerifyAuthorization(handleGetAlbumByID)).Methods("GET")
	r.HandleFunc("/albums", auth.VerifyAuthorization(handlePostAlbum)).Methods("POST")
	r.HandleFunc("/albums/{id}", auth.VerifyAuthorization(handlePatchAlbum)).Methods("PATCH")
	r.HandleFunc("/albums/{id}", auth.VerifyAuthorization(handleDeleteAlbum)).Methods("DELETE")
}
