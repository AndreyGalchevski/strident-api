package auth

import (
	"github.com/gorilla/mux"
)

func InitAuthRouter(router *mux.Router) {
	router.HandleFunc("/auth/login", handlePostLogin).Methods("POST")
	router.HandleFunc("/auth/verify", handleGetVerify).Methods("GET")
}
