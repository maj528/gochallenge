package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// Customers variable will load all customers and assign a variable (which will later be appended)
var Customers = LoadSamplesCsv()

// LoadSamplesCsv loads the sample csv file and returns the a Customer struct
func LoadSamplesCsv() []Customer {
	var customers []Customer
	csvFile, _ := os.Open("customers.csv")
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))
	// skips first line of CSV (headers)
	reader.Read()
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		customers = append(customers, Customer{
			ID:           line[0],
			Name:         line[1],
			Email:        line[2],
			Text:         line[3],
			CreationTime: line[4],
		})
	}

	return customers
}

// GetCustomers handles the "/customers" request
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Customers)
}
