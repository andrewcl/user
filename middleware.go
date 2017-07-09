package main

import (
	"log"
	"net/http"
	"time"
	"fmt"
)

/// validateError checks a given error, logging and panicking in case of valid error.
func validateError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func logMessage(message string) {
	log.Printf(fmt.Sprintf("%v - %v", time.Now(), message))
}

func timeLogHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)

		logMessage(fmt.Sprintf("Time Elapsed for %v request: %v",
			r.Method, time.Since(start)))
	}
	return http.HandlerFunc(fn)
}

func recoveryHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic Occurred at")
				logMessage(fmt.Sprintf("Panic occurred: %+v", err))
				http.Error(w, http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
