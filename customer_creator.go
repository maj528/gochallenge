package main

import (
	"encoding/json"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

// CreateCustomer - the r.body.Decode sets up the variables - could possibly set a bad response if fields are not fully filled in here
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customer Customer
	customer.CreationTime = time.Now().Format(time.RFC3339)
	_ = json.NewDecoder(r.Body).Decode(&customer)
	customer.ID = uuid.NewV4().String()
	if customer.Name == "" || customer.Email == "" || customer.Text == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Requests did not include required info"))
	} else {
		json.NewEncoder(w).Encode(customer)
		Customers = append(Customers, customer)
	}
}
