package pg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PGConnection struct {
	conn   *sqlx.DB
	config DBConfig
}

func (pg *PGConnection) NewDBConnection(config DBConfig) (*PGConnection, error) {
	connInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database,
	)
	conn, err := sqlx.Open("postgres", connInfo)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return &PGConnection{
		conn:   conn,
		config: config,
	}, nil
}

func (pg *PGConnection) Close() error {
	return pg.conn.Close()
}

func (pg *PGConnection) GetPGConnection() *sqlx.DB {
	return pg.conn
}
