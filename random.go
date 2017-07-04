package main

import (
	"database/sql"

	rd "github.com/Pallinder/go-randomdata"
)

func addFakeUsers(db *sql.DB, count int) {
	for i := 0; i < count; i++ {
		addUser(db, User{FirstName: rd.FirstName(rd.RandomGender),
			LastName: rd.LastName(),
			Email:    rd.Email()})
	}
}
