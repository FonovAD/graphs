package model

type Pagination struct {
	Limit  int64 `db:"limit"`
	Offset int64 `db:"offset"`
}
