package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/customers", GetCustomers).Methods("GET")
	router.HandleFunc("/customers", CreateCustomer).Methods("POST")
	router.HandleFunc("/", GetMajid).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
