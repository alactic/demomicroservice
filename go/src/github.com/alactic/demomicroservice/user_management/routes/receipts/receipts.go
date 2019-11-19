package receipts

import (
	"github.com/alactic/mygoproject/controllers/receipts"
	"github.com/alactic/mygoproject/middlewares/authentication"
	"github.com/gorilla/mux"
)

func Receipts(router *mux.Router) {
	router.HandleFunc("/receipt", authentication.AuthMiddleware(receipts.CreateReceiptEndpoint)).Methods("POST")
	router.HandleFunc("/receipt/{id}", authentication.AuthMiddleware(receipts.GetReceiptEndpoint)).Methods("GET")
}
