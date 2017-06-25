package main

import (
	"net/http"
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"
)

func createPackageRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users", searchUserNames).Methods("GET")
	return r
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

}

func createUser(w http.ResponseWriter, r *http.Request) {
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	email := r.FormValue("email")

	user, err := NewUser(firstName, lastName, email)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}

	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusOK)
}

func searchUserNames(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Println("called function searchUserNames")

	firstLetter := vars["first_character"]
	fmt.Println(firstLetter)

	user := User{
		FirstName: "Andrew",
		LastName:  "Liu",
		Email:     "andrew@andrewcl.com",
	}

	json.NewEncoder(w).Encode(user)
}