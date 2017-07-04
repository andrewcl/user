package main

import (
	"encoding/json"
	"net/http"
)

func handlePostUsersRequest(w http.ResponseWriter, user User) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(user)
}

func handleGetUsersRequest(w http.ResponseWriter, users []User) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(users)
}

/// handleBadRequest returns a json response with a 400 error code.
/// Intended to invoked when invalid parameters are sent
func handleBadRequest(w http.ResponseWriter) {
	http.Error(w,
		"Bad Request. The value for one of the URL parameters was invalid",
		http.StatusBadRequest)
}

/// handleNotFoundRequest handles misdirected or unavailable routes
func handleNotFoundRequest(w http.ResponseWriter) {
	http.Error(w,
		"Not Found: No API method associated with the URL path of the request",
		http.StatusNotFound)
}
