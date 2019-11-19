package customers

import (
	"github.com/alactic/mygoproject/controllers/customers"
	"github.com/alactic/mygoproject/middlewares/authentication"
	"github.com/gorilla/mux"
)

func Customers(router *mux.Router) {
	router.HandleFunc("/", customers.GetIndexEndpoint).Methods("GET")
	router.HandleFunc("/customer", customers.CreateCustomerEndpoint).Methods("POST")
	router.HandleFunc("/customer/{id}", authentication.AuthMiddleware(customers.GetCustomerEndpoint)).Methods("GET")
	router.HandleFunc("/customers", customers.GetCustomersEndpoint).Methods("GET")
	router.HandleFunc("/uploads", authentication.AuthMiddleware(customers.UploadFile)).Methods("POST")
	// router.HandleFunc("/readFiles", customers.ReadFile).Methods("GET")
}
