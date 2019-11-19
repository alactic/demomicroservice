package products

import (
	"github.com/alactic/mygoproject/controllers/products"
	"github.com/alactic/mygoproject/middlewares/authentication"
	"github.com/gorilla/mux"
)

func Products(router *mux.Router) {
	router.HandleFunc("/product", authentication.AuthMiddleware(products.CreateProductEndpoint)).Methods("POST")
	router.HandleFunc("/product/{id}", authentication.AuthMiddleware(products.GetProductEndpoint)).Methods("GET")
	router.HandleFunc("/products", authentication.AuthMiddleware(products.GetProductsEndpoint)).Methods("GET")
}
