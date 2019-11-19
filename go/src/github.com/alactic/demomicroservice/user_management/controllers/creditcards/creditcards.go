package creditcards

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alactic/mygoproject/models/creditcard"
	"github.com/alactic/mygoproject/models/customers"
	"github.com/alactic/mygoproject/models/product"
	"github.com/alactic/mygoproject/models/receipt"
	"github.com/alactic/mygoproject/utils/connection"
	"github.com/gorilla/mux"
	"gopkg.in/couchbase/gocb.v1"
)

type CreditCard = creditcard.CreditCard
type Customer = customers.Customer
type Product = product.Product
type Receipt = receipt.Receipt

var bucket *gocb.Bucket = connection.Connection()

//router.HandleFunc("/customer/creditcard/{id}", AddCreditCardEndpoint).Methods("PUT")
func AddCreditCardEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	routeParams := mux.Vars(request)
	var creditcard CreditCard
	_ = json.NewDecoder(request.Body).Decode(&creditcard)
	fmt.Println(routeParams["id"])
	fmt.Println(creditcard)
	_, err := bucket.MutateIn(routeParams["id"], 0, 0).ArrayAppend("creditcards", creditcard, true).Execute()
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{"message: "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(creditcard)
}

//router.HandleFunc("/customer/creditcard/{id}", AddCreditCardEndpoint).Methods("GET")
func GetCreditCardEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	routeParams := mux.Vars(request)
	fragment, err := bucket.LookupIn(routeParams["id"]).Get("creditcards").Execute()
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{"message: "` + err.Error() + `"}`))
	}
	var creditcard []CreditCard
	fragment.Content("creditcard", &creditcard)
	json.NewEncoder(response).Encode(creditcard)
}
