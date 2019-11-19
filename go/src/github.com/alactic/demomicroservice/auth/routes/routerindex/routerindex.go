package routerindex

import (
	"github.com/alactic/demomicroservice/auth/routes/auth"
	"github.com/gorilla/mux"
)

func Routerindex(router *mux.Router) {
	auth.Auth(router)
}
