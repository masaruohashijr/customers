package main

import (
	"net/http"
	"phonenumbers_backend/database"
	"phonenumbers_backend/features/customer"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDatabase()
	router := mux.NewRouter()
	// Customer
	router.HandleFunc("/customer", customer.GetAll).Methods("GET", "OPTIONS")
	router.HandleFunc("/customer/{id}", customer.Delete).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/customer/{id}", customer.Get).Methods("GET", "OPTIONS")
	router.HandleFunc("/customer", customer.Post).Methods("POST", "OPTIONS")
	router.HandleFunc("/customer", customer.Put).Methods("PUT", "OPTIONS")
	http.Handle("/", router)
	addr := ":5000"
	println("Phone Numbers")
	http.ListenAndServe(addr, router)
}
