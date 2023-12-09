package model

import (
	"fmt"
	"regexp"
)

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Login    string `json:"login"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Email == "" {
		return fmt.Errorf("email required")
	}
	if !isEmailValid(r.Email) {
		return fmt.Errorf("email is not valid")
	}
	if r.Password == "" {
		return fmt.Errorf("password required")
	}
	if r.Login == "" {
		return fmt.Errorf("login required")
	}
	return nil
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
