package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	UserPostParamFirstName = "first_name"
	UserPostParamLastName  = "last_name"
	UserPostParamEmail     = "email"

	UserGetParamStartingChar = "starting_char"
)

func createPackageRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handleRoot).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users", searchUserNames).Methods("GET")
	return r
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handleNotFoundRequest(w)
	}

	//TODO: serve testing / redirect webpage
}

func createUser(w http.ResponseWriter, r *http.Request) {
	firstName := r.FormValue(UserPostParamFirstName)
	lastName := r.FormValue(UserPostParamLastName)
	email := r.FormValue(UserPostParamEmail)

	user, err := NewUser(firstName, lastName, email)
	if err != nil {
		handleBadRequest(w)
		return
	}

	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusOK)
}

func searchUserNames(w http.ResponseWriter, r *http.Request) {
	lastNameChar := r.URL.Query().Get(UserGetParamStartingChar)

	if len(lastNameChar) < 1 {
		handleBadRequest(w)
		return
	}

	if len(lastNameChar) > 1 {
		lastNameChar = string([]rune(lastNameChar)[0])
	}

	accessDB := appDB

	users := retrieveUsersByLastName(accessDB, lastNameChar)
	//		json.NewEncoder(w).Encode(users)
	handleGetUsersRequest(w, users)
}
