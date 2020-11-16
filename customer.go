package main

// Customer is the 'struct' - like a class
type Customer struct {
	ID           string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name         string `json:"name,omitempty"`
	Email        string `json:"email,omitempty"`
	Text         string `json:"text,omitempty"`
	CreationTime string `json:"creationtime,omitempty"`
}
