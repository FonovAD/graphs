package controller

import "errors"

var ErrShortFirstname = errors.New("firstname is too short")

var ErrShortLastname = errors.New("lastname is too short")

var ErrShortPassword = errors.New("password is too short")
