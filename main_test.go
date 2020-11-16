package main

import (
	"net/http"
	"os"
	"testing"
)

func TestCSV(t *testing.T) {
	fi, _ := os.Stat("customers.csv")
	// get the size
	size := fi.Size()
	if size == 0 {
		t.Errorf("CSV file is empty - please download or rename 'customers.csv'")
	}

}

func TestGo(t *testing.T) {
	if 1 != 1 {
		t.Errorf("go has broken")
	}
}

func TestAPI(t *testing.T) {
	resp := GetCustomers
	if resp == nil {
		t.Errorf("API connection not working")
	}
}

func TestInternetConnection(t *testing.T) {
	_, err := http.Get("https://www.bbc.com/")
	if err != nil {
		t.Errorf("Internet is not connected.")
	}
}

func TestCustomer(t *testing.T) {
	testCustomer := Customer{}
	testCustomer.Name = "Mr Test"
	testCustomer.Email = "test@test"
	testCustomer.ID = "1000"
	testCustomer.CreationTime = "Today"
	if testCustomer.Text != "" {
		t.Errorf("Customer struct is created but auto creating text")
	}
	testCustomer.Text = "Lorem ipsum"
	testCustomer2 := Customer{}
	testCustomer2.Name = "Mr Test"
	testCustomer2.Email = "test@test"
	testCustomer2.ID = "1000"
	testCustomer2.CreationTime = "Today"
	testCustomer2.Text = "Lorem ipsum"
	if testCustomer != testCustomer2 {
		t.Errorf("Customer structs are unequal")
	}
}
