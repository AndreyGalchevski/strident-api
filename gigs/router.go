package gigs

import (
	"github.com/AndreyGalchevski/strident-api/auth"
	"github.com/gorilla/mux"
)

func InitGigsRouter(r *mux.Router) {
	r.HandleFunc("/gigs", handleGetGigs).Methods("GET")

	r.HandleFunc("/gigs/{id}", auth.VerifyAuthorization(handleGetGigByID)).Methods("GET")
	r.HandleFunc("/gigs", auth.VerifyAuthorization(handlePostGig)).Methods("POST")
	r.HandleFunc("/gigs/{id}", auth.VerifyAuthorization(handlePatchGig)).Methods("PATCH")
	r.HandleFunc("/gigs/{id}", auth.VerifyAuthorization(handleDeleteGig)).Methods("DELETE")
}
