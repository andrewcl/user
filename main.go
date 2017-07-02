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

	addUser(app.DB, &User{"Andrew", "Liu", "andrew@andrewcl.com"})
	addUser(app.DB, &User{"Leslie", "Chang", "leslie@chang.com"})
	addUser(app.DB, &User{"Jonathan", "Chang", "Jonathan@chang.com"})

	router := mux.NewRouter()
	router.HandleFunc("/users", app.createUser).Methods("POST")
	router.HandleFunc("/users", app.searchLastNames).Methods("GET")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(APPPORT, nil))
}
