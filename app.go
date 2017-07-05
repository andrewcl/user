package main

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
)

const (
	UserPostParamFirstName = "first_name"
	UserPostParamLastName  = "last_name"
	UserPostParamEmail     = "email"

	UserGetParamStartingChar = "starting_char"
	UserGetParamPage         = "page"
	UserGetParamPerPage      = "per_page"
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

	addUser(a.DB, *user)
	if err != nil {
		handleBadRequest(w)
		return
	}

	handlePostUsersRequest(w, *user)
}

func (a *App) searchLastNames(w http.ResponseWriter, r *http.Request) {
	lastNameChar := r.URL.Query().Get(UserGetParamStartingChar)
	page := convertIntQueryParameter(r, UserGetParamPage, 0)
	perPage := convertIntQueryParameter(r, UserGetParamPerPage, 10)

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

	users, err := retrieveLastNameUsers(a.DB, lastNameChar, page, perPage)
	if err != nil {
		handleBadRequest(w)
		return
	}
	handleGetUsersRequest(w, users)
}

func convertIntQueryParameter(r *http.Request, varName string, defaultVal int) int {
	stringParam := r.URL.Query().Get(varName)
	if stringParam == "" {
		return defaultVal
	}

	if intParam, err := strconv.Atoi(stringParam); err == nil {
		return intParam
	}

	return defaultVal
}
