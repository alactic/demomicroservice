package auth

import (
	"github.com/alactic/demomicroservice/auth/controllers/auth"
	"github.com/gorilla/mux"
)

func Auth(router *mux.Router) {
	router.HandleFunc("/", auth.TestAuth).Methods("GET")
	router.HandleFunc("/login", auth.LoginEndpoint).Methods("POST")
	router.HandleFunc("/signup", auth.SignupEndpoint).Methods("POST")
}
