package main

import (
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

func main() {
	rootURI := fmt.Sprintf("%v:%v@%v/", DBUSERNAME, DBPASSWORD, DBADDRESS)
	app := App{DB: createDatabase(DRIVERNAME, rootURI, DBNAME, DBUSERSTABLE)}
	defer app.DB.Close()

	addUser(app.DB, User{"Andrew", "Liu", "andrew@andrewcl.com"})
	addUser(app.DB, User{"Leslie", "Chang", "leslie@chang.com"})
	addUser(app.DB, User{"Jonathan", "Chang", "Jonathan@chang.com"})
	addFakeUsers(app.DB, 3000)

	http.Handle("/", http.HandlerFunc(app.routeApiRouteRoot))
	http.Handle("/users", timeLogHandler(recoveryHandler(http.HandlerFunc(app.routeApiRouteUser))))
	log.Fatal(http.ListenAndServe(APPPORT, nil))
}
