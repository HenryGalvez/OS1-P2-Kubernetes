package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type casee struct {
	Name			string	`json:"name"`
	Location		string	`json:"location"`
	Age				int		`json:"age"`
	Infectedtype 	string	`json:"infectedtype"`
	State 			string	`json:"state"`
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GO API!")
}

func createCase(w http.ResponseWriter, r *http.Request) {
	var newcase casee
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Request Invalido")
	}
	json.Unmarshal(reqBody, &newcase)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newcase)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/cases", createCase).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))
}
