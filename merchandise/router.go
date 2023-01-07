package merchandise

import (
	"github.com/AndreyGalchevski/strident-api/auth"
	"github.com/gorilla/mux"
)

func InitMerchandiseRouter(r *mux.Router) {
	r.HandleFunc("/merchandise", handleGetMerchandise).Methods("GET")

	r.HandleFunc("/merchandise/{id}", auth.VerifyAuthorization(handleGetMerchandiseByID)).Methods("GET")
	r.HandleFunc("/merchandise", auth.VerifyAuthorization(handlePostMerchandise)).Methods("POST")
	r.HandleFunc("/merchandise/{id}", auth.VerifyAuthorization(handlePatchMerchandise)).Methods("PATCH")
	r.HandleFunc("/merchandise/{id}", auth.VerifyAuthorization(handleDeleteMerchandise)).Methods("DELETE")
}
