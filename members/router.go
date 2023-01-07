package members

import (
	"github.com/AndreyGalchevski/strident-api/auth"
	"github.com/gorilla/mux"
)

func InitMembersRouter(r *mux.Router) {
	r.HandleFunc("/members", handleGetMembers).Methods("GET")

	r.HandleFunc("/members/{id}", auth.VerifyAuthorization(handleGetMemberByID)).Methods("GET")
	r.HandleFunc("/members", auth.VerifyAuthorization(handlePostMember)).Methods("POST")
	r.HandleFunc("/members/{id}", auth.VerifyAuthorization(handlePatchMember)).Methods("PATCH")
	r.HandleFunc("/members/{id}", auth.VerifyAuthorization(handleDeleteMember)).Methods("DELETE")
}
