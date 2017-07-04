package main

import (
	"database/sql"
	"net/http"
	"strings"
)

const (
	UserPostParamFirstName = "first_name"
	UserPostParamLastName  = "last_name"
	UserPostParamEmail     = "email"

	UserGetParamStartingChar = "starting_char"
)

type App struct {
	DB *sql.DB
}

func (a *App) createUser(w http.ResponseWriter, r *http.Request) {
	firstName := r.FormValue(UserPostParamFirstName)
	lastName := r.FormValue(UserPostParamLastName)
	email := r.FormValue(UserPostParamEmail)

	user, err := NewUser(firstName, lastName, email)
	if err != nil {
		handleBadRequest(w)
		return
	}

	addUser(a.DB, user)
	if err != nil {
		handleBadRequest(w)
		return
	}

	handlePostUsersRequest(w)
}

func (a *App) searchLastNames(w http.ResponseWriter, r *http.Request) {
	lastNameChar := r.URL.Query().Get(UserGetParamStartingChar)

	if len(lastNameChar) < 1 {
		handleBadRequest(w)
		return
	}

	// truncate strings greater than one character
	if len(lastNameChar) > 1 {
		lastNameChar = string([]rune(lastNameChar)[0])
	}

	// convert string to lowercase
	lastNameChar = strings.ToLower(lastNameChar)

	users, err := retrieveLastNameUsers(a.DB, lastNameChar)
	if err != nil {
		handleBadRequest(w)
		return
	}
	handleGetUsersRequest(w, users)
}
