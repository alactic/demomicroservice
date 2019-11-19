package staff

import (
	"github.com/alactic/mygoproject/controllers/staff"
	"github.com/alactic/mygoproject/middlewares/authentication"
	"github.com/gorilla/mux"
)

func Staff(router *mux.Router) {
	router.HandleFunc("/staff", staff.CreateStaffEndpoint).Methods("POST")
	router.HandleFunc("/staff/{id}", authentication.AuthMiddleware(staff.GetStaffEndpoint)).Methods("GET")
	router.HandleFunc("/staff", staff.GetAllStaffEndpoint).Methods("GET")
	router.HandleFunc("/uploads", authentication.AuthMiddleware(staff.UploadFile)).Methods("POST")
	// router.HandleFunc("/readFiles", staff.ReadFile).Methods("GET")
}
