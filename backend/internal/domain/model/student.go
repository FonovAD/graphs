package model

type Student struct {
	ID      int64 `db:"student_id"`
	UserID  int64 `db:"usersid"`
	GroupID int64 `db:"groups_id"`
}
