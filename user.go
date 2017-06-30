package main

import "errors"

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type Users []User

func NewUser(firstName, lastName, email string) (*User, error) {
	if firstName == "" {
		return nil, errors.New("Invalid First Name")
	}
	if lastName == "" {
		return nil, errors.New("Invalid Last Name")
	}
	if email == "" {
		//TODO: add validation for email.
		// Consider Regex or external Golang library
		return nil, errors.New("Invalid Email")
	}
	return &User{FirstName: firstName, LastName: lastName, Email: email},
		nil
}
