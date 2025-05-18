package usecase

import "errors"

var ErrShortFirstname = errors.New("firstname is too short")

var ErrLongFirstname = errors.New("firstname is too long")

var ErrShortLastname = errors.New("lastname is too short")

var ErrLongLastname = errors.New("lastname is too long")

var ErrShortFathername = errors.New("fathername is too short")

var ErrLongFathername = errors.New("fathername is too long")

var ErrShortPassword = errors.New("password is too short")

var ErrLongPassword = errors.New("password is too long")

var ErrShortEmail = errors.New("email is too short")

var ErrLongEmail = errors.New("email is too long")

var ErrParseToken = errors.New("parse token error")

var ErrNoPermissions = errors.New("no permissions")
