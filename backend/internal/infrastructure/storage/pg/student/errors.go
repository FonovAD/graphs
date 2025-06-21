package storage

import "errors"

var (
	ErrTimeExceeded      = errors.New("time exceeded")
	ErrAnswerAlreadySent = errors.New("answer already sent")
)
