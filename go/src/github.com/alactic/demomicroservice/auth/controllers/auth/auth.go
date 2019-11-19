package auth

import (
	"encoding/json"
	"net/http"

	"github.com/alactic/demomicroservice/user_management/models/customers"
	"github.com/gorilla/mux"
	"github.com/alactic/demomicroservice/auth/models/auth"
	hashed "github.com/alactic/demomicroservice/library/utils/hash"
//	jwtFile "github.com/alactic/demomicroservice/library/utils/jwt"
	"github.com/alactic/demomicroservice/library/utils/connection"
	"gopkg.in/couchbase/gocb.v1"
	uuid "github.com/satori/go.uuid"
)

type Customer = customers.Customer
var bucket *gocb.Bucket = connection.Connection()

// TESTING Auth endpoint

func TestAuth (response http.ResponseWriter, request *http.Request) {
	response.Write([]byte(`{"message": "auth endpoint working"}`))
}

// Login endpoint for authentication
func LoginEndpoint(response http.ResponseWriter, request *http.Request) {
	// jwtFile.DecodeJWT(request.Header["Authorization"][0])
	response.Header().Set("Content-Type", "application/json")
	var auth auth.Auth
	_ = json.NewDecoder(request.Body).Decode(&auth)
	hashed.Hash(auth.Password)
	hashed.CompareHashValue("password", "hashpassword")
	routerParams := mux.Vars(request)
	var customer Customer
	_, err := bucket.Get(routerParams["id"], &customer)
	if err != nil {
		if err.Error() == "key not found" {
			response.Write([]byte(`{"data": "{}"}`))
			return
		}
		response.WriteHeader(500)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(customer)
	// details := make(map[string]string)
	// details["firstname"] = "echezona"
	// details["lastname"] = "okafor"
	// details["email"] = "okaforechezona1992@gmail.com"
	// var token = jwtFile.GenerateJWT(details)
	// json.NewEncoder(response).Encode(token)
}

func SignupEndpoint(response http.ResponseWriter, request *http.Request) {
	// jwtFile.DecodeJWT(request.Header["Authorization"][0])
	// response.Header().Set("Content-Type", "application/json")
	// var auth auth.Auth
	// _ = json.NewDecoder(request.Body).Decode(&auth)
	// hashed.Hash(auth.Password)
	// details := make(map[string]string)
	// details["firstname"] = "echezona"
	// details["lastname"] = "okafor"
	// details["email"] = "okaforechezona1992@gmail.com"
	// var token = jwtFile.GenerateJWT(details)
	// details["token"] = token


	// response.Header().Set("Access-Control-Allow-Origin", "*")
	// response.Header().Set("Content-Type", "application/json")
	var customer Customer
	_ = json.NewDecoder(request.Body).Decode(&customer)
	id := uuid.Must(uuid.NewV4()).String()
	customer.Type = "staff"
	customer.Id = id
	_, err := bucket.Insert(id, customer, 0)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{"message": "` + err.Error() + `" }`))
		return
	}
	 response.Write([]byte(`{ "id": "` + id + `"}`))


	// json.NewEncoder(response).Encode(details)
}
