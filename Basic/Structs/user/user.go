package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstname string
	lastname  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) Admin {

	return Admin{
		email:    email,
		password: password,
		User: User{
			firstname: "Admin",
			lastname:  "Admin",
			birthDate: "01.01.1970",
			createdAt: time.Now(),
		},
	}
}

func NewUser(firstname, lastname, birthDate string) (*User, error) {
	if firstname == "" || lastname == "" || birthDate == "" {
		return nil, errors.New("Firstname, lastname or birth date is empty")
	}
	return &User{
		firstname: firstname,
		lastname:  lastname,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func (u User) OutputUser() {
	fmt.Println("Name: ", u.firstname, u.lastname, "\nBirth date: ", u.birthDate, "\nCreated at: ", u.createdAt)
}

func (u *User) ClearUserData() {
	u.firstname = ""
	u.lastname = ""
}
