package main

import (
	"log"
	"net/http"
	"context"

	
	proto "github.com/alactic/demosample/userservice/proto/user"
	"github.com/micro/go-micro"
	"github.com/gorilla/handlers"
)


func main() {
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	service := micro.NewService(micro.Name("clientuser"))
	service.Init()

	client := proto.NewUserService("user", service.Client())

	r, err := client.UserDetails(context.Background(), 23)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r)

	log.Fatal(http.ListenAndServe("0.0.0.0:8806", handlers.CORS(headers, methods, origins)(router)))
}