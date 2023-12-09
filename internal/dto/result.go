package dto

import "time"

type Result struct {
	ID        int64
	Start     time.Time
	End       time.Time
	Grade     int64
	StudentID int64
	TestID    int64
}
