package products

import (
	"encoding/json"
	"net/http"

	"github.com/alactic/mygoproject/models/creditcard"
	"github.com/alactic/mygoproject/models/customers"
	"github.com/alactic/mygoproject/models/product"
	"github.com/alactic/mygoproject/models/receipt"
	"github.com/alactic/mygoproject/utils/connection"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/couchbase/gocb.v1"
)

type CreditCard = creditcard.CreditCard
type Customer = customers.Customer
type Product = product.Product
type Receipt = receipt.Receipt

var bucket *gocb.Bucket = connection.Connection()

//router.HandleFunc("/product", CreateProductEndpoint).Methods("POST")
func CreateProductEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var product Product
	_ = json.NewDecoder(request.Body).Decode(&product)
	id := uuid.Must(uuid.NewV4()).String()
	product.Type = "product"
	product.Id = id
	_, err := bucket.Insert(id, product, 0)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{"message": "` + err.Error() + `" }`))
		return
	}
	response.Write([]byte(`{ "id": "` + id + `"}`))
}

//router.HandleFunc("/product/{id}", GetProductEndpoint).Methods("GET")
func GetProductEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	routerParams := mux.Vars(request)
	var product Product
	product.Id = routerParams["id"]
	_, err := bucket.Get(routerParams["id"], &product)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(product)
}

//router.HandleFunc("/product", GetProductsEndpoint).Methods("GET")
func GetProductsEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var products []Product
	query := gocb.NewN1qlQuery("SELECT META().id, " + bucket.Name() + ".* FROM " + bucket.Name() + " WHERE type = 'product'")
	rows, err := bucket.ExecuteN1qlQuery(query, nil)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	var row Product
	for rows.Next(&row) {
		products = append(products, row)
	}
	json.NewEncoder(response).Encode(products)
}
