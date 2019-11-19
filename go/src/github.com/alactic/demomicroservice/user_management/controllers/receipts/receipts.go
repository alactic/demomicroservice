package receipts

import (
	"encoding/json"
	"net/http"

	"github.com/alactic/mygoproject/models/receipt"
	"github.com/alactic/mygoproject/utils/connection"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/couchbase/gocb.v1"
)

type Receipt = receipt.Receipt

var bucket *gocb.Bucket = connection.Connection()

//router.HandleFunc("/receipt", CreateReceiptEndpoint).Methods("POST")
func CreateReceiptEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var receipt Receipt
	_ = json.NewDecoder(request.Body).Decode(&receipt)
	productsIds := make([]string, len(receipt.Product))
	for i, product := range receipt.Product {
		productsIds[i] = product.Id
	}
	query := gocb.NewN1qlQuery("SELECT (SELECT VALUE {META(customer).id, customer.firstname, customer.lastname, customer.type} FROM " +
		bucket.Name() + " As customer USE KEYS $2)[0] AS customer, products FROM " + bucket.Name() +
		" AS receipt USE KEYS $2 LET products =(SELECT META(product).id, product.price, price.type FROM " +
		bucket.Name() + " AS product USE KEYS $1)")

	var params []interface{}
	params = append(params, productsIds)
	params = append(params, receipt.Customer.Id)
	rows, err := bucket.ExecuteN1qlQuery(query, params)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	receipt = Receipt{}
	err = rows.One(&receipt)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	receipt.Type = "receipt"
	for _, product := range receipt.Product {
		receipt.Total += product.Price
	}
	id := uuid.Must(uuid.NewV4()).String()
	bucket.Insert(id, receipt, 0)
	response.Write([]byte(`{"id" : "` + id + `" }`))

}

//router.HandleFunc("/receipt/{id}", GetReceiptEndpoint).Methods("GET")
func GetReceiptEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	routeParams := mux.Vars(request)
	query := gocb.NewN1qlQuery("SELECT META(receipt).id, receipt.* FROM " + bucket.Name() + " AS receipts  WHERE receipt.type = 'receipt' AND receipts.customer.id = $1")
	var params []interface{}
	params = append(params, routeParams["id"])
	rows, err := bucket.ExecuteN1qlQuery(query, params)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	var receipts []Receipt
	var row Receipt
	for rows.Next(&row) {
		receipts = append(receipts, row)
	}
	json.NewEncoder(response).Encode(receipts)
}
