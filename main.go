package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	DRIVERNAME   = "mysql"
	DBADDRESS    = "tcp(localhost:3306)"
	DBNAME       = "uservoice"
	DBUSERSTABLE = "Users"
	DBUSERNAME   = "root"
	DBPASSWORD   = "root"

	APPPORT = ":8080"
)

func main() {
	rootURI := fmt.Sprintf("%v:%v@%v/", DBUSERNAME, DBPASSWORD, DBADDRESS)
	app := App{DB: createDatabase(DRIVERNAME, rootURI, DBNAME, DBUSERSTABLE)}
	defer app.DB.Close()

	router := mux.NewRouter()
	router.HandleFunc("/users", app.createUser).Methods("POST")
	router.HandleFunc("/users", app.searchLastNames).Methods("GET")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(APPPORT, nil))
}
