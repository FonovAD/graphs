package model

import (
	"fmt"
	"time"
)

type Lab struct {
	ID               int64     `db:"lab_id" json:"id"`
	Name             string    `db:"name" json:"name"`
	Description      string    `db:"description" json:"description"`
	Duration         string    `db:"duration" json:"duration"`
	RegistrationDate time.Time `db:"registration_date" json:"registrationDate"`
	TeacherID        int64     `db:"teacher_id" json:"teacherId"`
	TeacherFIO       string    `db:"teacher_fio" json:"teacherFio"`
}

func (l *Lab) SetDurationMinutes(minutes int64) {
	h := minutes / 60
	m := minutes % 60
	l.Duration = fmt.Sprintf("%02d:%02d:00", h, m)
}

func (l *Lab) GetDurationMinutes() int64 {
	t, err := time.Parse("15:04:05", l.Duration)
	if err != nil {
		return 0
	}
	return int64(t.Hour()*60 + t.Minute())
}
