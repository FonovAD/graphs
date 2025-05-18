package model

import "time"

type User struct {
	ID               int64     `db:"usersid"`
	DateRegistration time.Time `db:"date_registration"`
	Email            string    `db:"email"`
	Password         string    `db:"password"`
	FirstName        string    `db:"first_name"`
	LastName         string    `db:"last_name"`
	FatherName       string    `db:"father_name"`
	Role             string    `db:"role"`
	PasswordSalt     string    `db:"passwordsalt"`
}
