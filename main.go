package main

import (
	"log"
	"net/http"
)

func main() {
	router := createPackageRouter()

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

