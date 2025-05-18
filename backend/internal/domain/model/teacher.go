package model

type Teacher struct {
	ID     int64 `db:"teacherid"`
	UserID int64 `db:"usersid"`
}
