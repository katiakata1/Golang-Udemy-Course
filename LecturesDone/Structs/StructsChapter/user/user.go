package user

import (
	"errors"
	"fmt"
	"time"
)

// I am making User global because I want it to be accessible outside this package
// But if I want fields also be accessible, I need to change firstName to FirstName
// However, we don't need to this now.
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// embedded structs are used to build on the existing struct and its methods
// we can do User (field name) and User (struct User) but then we will have to access User methods like admin.User.OutputUserDetails()
// Also, we can use anonomous field and just call it User (struct name) and then you will be access method with admin.OutputUserDetails()
type Admin struct {
	email    string
	password string
	User
}

func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// Embedded constructor function
// If you use anonomous field, you still need to write User (repeating the name of the sturct): User (2nd User is a struct)
func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "____",
			createdAt: time.Now(),
		},
	}
}

// We call this function from main package so it would create a struct for us
func NewUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("firstname, lastname and birthdate are required")
	}
	// returning a pointer
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
