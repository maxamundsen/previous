package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Json endpoints for an API can be easily written with a handler.
// In order to parse json, you must specify some information about the output json
// via decorators.
// structs can be automatically generated from json using: https://mholt.github.io/json-to-go/

type person struct {
	FirstName string
	LastName  string
	Age       int
}

// This endpoint creates a person struct and returns encoded JSON
func apiTestHandler(w http.ResponseWriter, r *http.Request) {
	data := person{
		"John",
		"Doe",
		20,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// This will print information about the current Identity
// via session access
func apiUserHandler(w http.ResponseWriter, r *http.Request) {
	user := sessionStore.GetIdentityFromCtx(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Struct for serializing json from 3rd party API
// http://api.open-notify.org/astros.json
type astro struct {
	Message string `json:"message"`
	People  []struct {
		Name  string `json:"name"`
		Craft string `json:"craft"`
	} `json:"people"`
	Number int `json:"number"`
}

// This endpoint fetches a JSON response from a third party API,
// serializes it to an 'astro' struct, then
func apiClientFetchHandler(w http.ResponseWriter, r *http.Request) {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	req, reqErr := http.NewRequest(http.MethodGet, "http://api.open-notify.org/astros.json", nil)

	if reqErr != nil {
		log.Println(reqErr)
	}

	req.Header.Set("User-Agent", "Example-Api")

	res, resErr := client.Do(req)

	if resErr != nil {
		log.Println(resErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		log.Println(readErr)
	}

	jsonOutput := astro{}

	jsonErr := json.Unmarshal(body, &jsonOutput)

	if jsonErr != nil {
		log.Println(jsonErr)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonOutput)
}
