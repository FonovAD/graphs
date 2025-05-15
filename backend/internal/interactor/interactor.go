package interactor

import "github.com/jmoiron/sqlx"

type Interactor interface{}

type interactor struct {
	conn *sqlx.DB
}

func NewInteractor(conn *sqlx.DB) Interactor {
	return &interactor{conn: conn}
}
