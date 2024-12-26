package router

import (
	"github.com/gorilla/mux"
	"crud/controller"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/create", controller.Create).Methods("POST")
	r.HandleFunc("/api/get", controller.GetAll).Methods("GET")
	r.HandleFunc("/api/get/{id}", controller.Get).Methods("GET")
	r.HandleFunc("/api/delete", controller.DeleteAll).Methods("DELETE")
	r.HandleFunc("/api/delete/{id}", controller.Delete).Methods("DELETE")

	return r
}
