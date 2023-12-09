package model

import "fmt"

type AuthUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthUserResponse struct {
	Token string `json:"token"`
}

func (r *AuthUserRequest) Validate() error {
	if r.Email == "" {
		return fmt.Errorf("email required")
	}
	if !isEmailValid(r.Email) {
		return fmt.Errorf("email is not valid")
	}
	if r.Password == "" {
		return fmt.Errorf("password required")
	}
	return nil
}
