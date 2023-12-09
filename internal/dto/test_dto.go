package dto

import "time"

type Test struct {
	ID       int64
	Name     string
	Start    time.Time
	End      time.Time
	Interval time.Time
}
