package dto

import "time"

type Test struct {
	ID          int64
	Name        string
	Description string
	Start       time.Time
	End         time.Time
}
