package main

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

type Founder struct {
	Name    string `json:"name"`
	Age     uint32 `json:"age"`
	Email   string `json:"email"`
	Company string `json:"company"`
}

var founders []Founder

func greetingsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var founder Founder
	json.NewDecoder(r.Body).Decode(&founder)
	founder.Age = founder.Age * 2
	founders = append(founders, founder)
	json.NewEncoder(w).Encode(founders)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", greetingsHandler).Methods("GET")
	r.HandleFunc("/form", formHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
