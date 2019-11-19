package creditcards

import (
	"github.com/alactic/mygoproject/controllers/creditcards"
	"github.com/alactic/mygoproject/middlewares/authentication"
	"github.com/gorilla/mux"
)

func Creditcards(router *mux.Router) {
	router.HandleFunc("/customer/creditcard/{id}", authentication.AuthMiddleware(creditcards.AddCreditCardEndpoint)).Methods("PUT")
	router.HandleFunc("/customer/creditcard/{id}", authentication.AuthMiddleware(creditcards.AddCreditCardEndpoint)).Methods("GET")
}
