package dto

import "time"

type User struct {
	Id               int64
	DateRegistration time.Time
	Email            string
	Password         string
	FirstName        string
	LastName         string
	Role             string
	PasswordSalt     string
}
