package model

type Group struct {
	ID   int64  `db:"groups_id"`
	Name string `db:"groupsname"`
}
