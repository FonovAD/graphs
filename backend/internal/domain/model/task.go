package model

import "database/sql"

type Task struct {
	ID       int64          `db:"task_id" json:"taskID"`
	ModuleID int64          `db:"module_id" json:"moduleID"`
	Payload  string         `db:"payload"`
	Answer   sql.NullString `db:"answer"`
}
