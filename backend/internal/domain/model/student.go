package model

type Student struct {
	ID      int64 `db:"student_id"`
	UserID  int64 `db:"usersid"`
	GroupID int64 `db:"groups_id"`
}

type StudentWithInfo struct {
	UserID int64  `db:"usersid" json:"user_id"`
	FIO    string `db:"fio" json:"fio"`
}
