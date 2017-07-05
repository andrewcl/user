package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func createDatabase(driver, rootURI, dbName, tableName string) *sql.DB {
	db, err := sql.Open(driver, rootURI)
	validateError(err)

	// Create database in case it does not exist
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", dbName))
	validateError(err)
	db.Close()

	// Open database w/ new URI
	dbURI := rootURI + dbName
	db, err = sql.Open(driver, dbURI)
	validateError(err)

	// Drop Table name in anticipation to recreate
	_, err = db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %v", tableName))
	validateError(err)

	_, err = db.Exec(fmt.Sprintf(
		`CREATE TABLE %v
	     (
		ID int auto_increment
		   primary key,
		FirstName varchar(100) not null,
		LastName varchar(100) not null,
		Email varchar(100) not null
	     );`, tableName))
	validateError(err)

	return db
}

func addUser(db *sql.DB, u User) error {
	stmt := fmt.Sprintf("INSERT INTO %v (FirstName, LastName, Email) VALUES (?, ?, ?)", DBUSERSTABLE)
	_, err := db.Exec(stmt, u.FirstName, u.LastName, u.Email)
	return err
}

func retrieveLastNameUsers(db *sql.DB, lastNameString string, page, perPage int) ([]User, error) {
	stmt := "SELECT FirstName, LastName, Email " +
		"FROM " + DBUSERSTABLE +
		" WHERE LEFT (LOWER(LastName), 1) = ?" +
		" LIMIT ? OFFSET ?"
	rows, err := db.Query(stmt, lastNameString, perPage, page)

	if err != nil {
		return nil, err
	}

	users := make([]User, 0)
	for rows.Next() {
		var u User
		_ = rows.Scan(&u.FirstName, &u.LastName, &u.Email)
		users = append(users, u)
	}
	return users, nil
}

/// validateError checks a given error, logging and panicking in case of valid error.
func validateError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
