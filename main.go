package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
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

var appDB *sql.DB

func main() {
	rootURI := fmt.Sprintf("%v:%v@%v/",
		DBUSERNAME, DBPASSWORD, DBADDRESS)
	appDB := createDatabase(DRIVERNAME, rootURI, DBNAME, DBUSERSTABLE)
	addUser(appDB, "Andrew", "Liu", "andrew@andrewcl.com")

	router := createPackageRouter()
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(APPPORT, nil))
}
