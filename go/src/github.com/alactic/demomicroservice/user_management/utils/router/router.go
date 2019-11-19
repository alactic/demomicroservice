package router

import (
	"log"
	"net/http"
	// "time"

	"github.com/alactic/mygoproject/routes/routerindex"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Router() {
	// r := mux.NewRouter()
	// v1 := r.PathPrefix("/api").Subrouter()
	// routerindex.Routerindex(v1)
	// if err := http.ListenAndServe(":8000", r); err != nil {
	// 	log.Fatal("Serving error.")
	// }

	// r := mux.NewRouter()
	// v1 := r.PathPrefix("/api").Subrouter()
	// routerindex.Routerindex(v1)
	// srv := &http.Server{
	// 	Handler: r,
	// 	Addr:    "0.0.0.0:8800",
	// 	// Good practice: enforce timeouts for servers you create!
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }

	// log.Fatal(srv.ListenAndServe())

	router := mux.NewRouter()
	v1 := router.PathPrefix("/api").Subrouter()
	routerindex.Routerindex(v1)
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe("0.0.0.0:8800", handlers.CORS(headers, methods, origins)(router)))

	// Apply the CORS middleware to our top-level router, with the defaults.
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	w.Write([]byte(`{"message": "testing cors"}`))
}
