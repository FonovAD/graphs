package model

import "time"

type Lab struct {
	ID               int64         `db:"lab_id" json:"id"`
	Name             string        `db:"name" json:"name"`
	Description      string        `db:"description" json:"description"`
	Duration         time.Duration `db:"duration" json:"duration"`
	RegistrationDate time.Time     `db:"registration_date" json:"registrationDate"`
	TeacherID        int64         `db:"teacher_id" json:"teacherId"`
	TeacherFIO       string        `db:"teacher_fio" json:"teacherFio"`
}

func (l *Lab) SetDurationMinutes(durationMinutes int64) {
	l.Duration = time.Duration(durationMinutes) * time.Minute
}

func (l *Lab) GetDurationMinutes() int64 {
	return int64(l.Duration.Minutes())
}
