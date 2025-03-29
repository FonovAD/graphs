package controller

import "golang_graphs/backend/internal/models"

func ValidateCreateUser(request models.CreateUserRequest) error {
	if len(request.FirstName) < 1 {
		return ErrShortFirstname
	}

	if len(request.LastName) < 1 {
		return ErrShortLastname
	}

	if len(request.Password) < 8 {
		return ErrShortPassword
	}

	return nil
}
