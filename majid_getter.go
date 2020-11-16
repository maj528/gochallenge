//A little information about me in a Go-friendly format
package main

import (
	"encoding/json"
	"net/http"
)

type Applicant struct {
	Name       string
	About      string
	Experience string
	Contact    string
	Me         string
}

func GetMajid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	majid := Applicant{"Majid Fakhoury", "Thank you for checking out my coding challenge! I can't wait to hear your feedback regarding my challenges and to hopefully join the team! I am a linguistics nerd and enjoy learning new languages. I love animals and am expecting a Shiba Inu puppy soon. I have worked in multiple countries, languages, and environments and thrive when collaborating with others. I truly believe in the value of diversity and am  excited to be surrounded by likeminded, forward-thinking tech professionals. I always tackle new challenges and never want to stop learning. I am a passionate cook and would love to cook for you one day!", "My experience is primarily in Ruby and JavaScript, where I am adept at working with APIs. However, I wanted a challenge - so I did my best to complete this challenge entirely in Go and learned along the way.", "majidfakhoury@gmail.com", "img src='https://i.imgur.com/1eOY0GH.png"}
	json.NewEncoder(w).Encode(majid)
}
