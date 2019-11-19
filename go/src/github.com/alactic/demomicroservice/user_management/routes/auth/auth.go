package auth

import (
	"github.com/alactic/mygoproject/controllers/auth"
	"github.com/gorilla/mux"
)

func Auth(router *mux.Router) {
	router.HandleFunc("/login", auth.LoginEndpoint).Methods("POST")
}
