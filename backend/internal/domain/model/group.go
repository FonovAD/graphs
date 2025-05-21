package model

type Group struct {
	ID   int64  `db:"groups_id" json:"GroupID"`
	Name string `db:"groupsname" json:"Name"`
}
