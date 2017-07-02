package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func handleClientResponse(w http.ResponseWriter, httpcode int, response []byte, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpcode)
	fmt.Fprintf(w, "{error:%v, code:%v, response:%v}", err, httpcode, response)
}

func handlePostUsersRequest(w http.ResponseWriter) {
	handleClientResponse(w, http.StatusOK, nil, nil)
}

func handleGetUsersRequest(w http.ResponseWriter, users []User) {
	userJSON, err := json.Marshal(users)
	if err != nil {
		//TODO: determine if error code more appropriate works here
		handleBadRequest(w)
	}
	handleClientResponse(w, http.StatusOK, userJSON, nil)
}

/// handleBadRequest returns a json response with a 400 error code.
/// Intended to invoked when invalid parameters are sent
func handleBadRequest(w http.ResponseWriter) {
	handleClientResponse(w, http.StatusBadRequest, nil, errors.New(
		"Bad Request. The value for one of the URL parameters was invalid"))
}

func handleNotFoundRequest(w http.ResponseWriter) {
	handleClientResponse(w, http.StatusNotFound, nil, errors.New(
		"Not Found: No API method associated with the URL path of the request"))
}