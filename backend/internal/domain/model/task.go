package model

import "database/sql"

type Task struct {
	ID       int64          `db:"task_id" json:"taskID"`
	ModuleID int64          `db:"module_id" json:"moduleID"`
	Payload  string         `db:"payload"`
	Answer   sql.NullString `db:"answer"`
}

type TaskByModule struct {
	ID      int64  `db:"task_id" json:"taskID"`
	Payload string `db:"payload" json:"payload"`
}

type AssignedTaskByModule struct {
	ID      int64         `db:"task_id" json:"taskID"`
	Payload string        `db:"payload" json:"payload"`
	Score   sql.NullInt64 `db:"score"   json:"score"`
}

type UserTask struct {
	UserID int64 `db:"user_id" json:"userID"`
	TaskID int64 `db:"task_id" json:"taskID"`
}

type TaskType struct {
	TaskType string `db:"task_type"`
}
