package model

import "time"

type UserLab struct {
	UserLabID      int64     `db:"user_lab_id"`
	UserID         int64     `db:"user_id"`
	LabID          int64     `db:"lab_id"`
	AssignmentDate time.Time `db:"assignment_date"`
	StartTime      time.Time `db:"start_time"`
	TeacherID      int64     `db:"teacher_id"`
	Deadline       time.Time `db:"deadline"`
	Score          *int      `db:"score"`
}

type UserLabGroup struct {
	UserLabID      int64     `db:"user_lab_id"`
	UserID         int64     `db:"user_id"`
	LabID          int64     `db:"lab_id"`
	AssignmentDate time.Time `db:"assignment_date"`
	StartTime      time.Time `db:"start_time"`
	TeacherID      int64     `db:"teacher_id"`
	Deadline       time.Time `db:"deadline"`
	Score          *int      `db:"score"`
	GroupID        int64     `db:"groups_id"`
}

type UserLabWithInfo struct {
	LabID   int64   `db:"lab_id" json:"labId"`
	LabName string  `db:"lab_name" json:"labName"`
	Groups  []Group `db:"groups" json:"groups"`
}

type UserLabAnswer struct {
	UserLabID int64  `db:"user_lab_id"`
	UserID    int64  `db:"user_id"`
	LabID     int64  `db:"lab_id"`
	TaskID    int64  `db:"task_id"`
	Answer    string `db:"answer"`
	Score     int    `db:"score"`
}
