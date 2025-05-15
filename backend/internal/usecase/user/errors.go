package usecase

import "errors"

var ErrShortPassword = errors.New("password is too short")

var ErrLongPassword = errors.New("password is too long")

var ErrShortEmail = errors.New("email is too short")

var ErrLongEmail = errors.New("email is too long")
