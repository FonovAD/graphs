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

	if len(request.Email) < 4 {
		return ErrShortEmail
	}

	if len(request.Email) > 100 {
		return ErrLongEmail
	}

	return nil
}

func ValidateAuthUser(request models.AuthUserRequest) error {
	if len(request.Password) < 8 {
		return ErrShortPassword
	}

	if len(request.Password) > 50 {
		return ErrLongPassword
	}

	if len(request.Email) < 4 {
		return ErrShortEmail
	}

	if len(request.Email) > 100 {
		return ErrLongEmail
	}

	return nil
}
