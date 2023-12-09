package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var (
	UserNotFoundError = fmt.Errorf("user doesn't exist")
)

type database struct {
	client *sql.DB
}

func NewDatabase(config Config) (Database, error) {
	connInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database,
	)
	client, err := sql.Open("postgres", connInfo)

	if err != nil {
		return nil, err
	}

	return &database{client: client}, nil
}

//
//func (db *database) Init(ctx context.Context) error {
//	_, err := db.client.ExecContext(ctx, initRequest)
//	return err
//}

//func (db *database) CreateUser(ctx context.Context, email, hash, login string) (int64, error) {
//	var id int64
//	row := db.client.QueryRowContext(ctx, createUser, email, hash, login)
//
//	err := row.Scan(&id)
//
//	if err != nil {
//		return 0, fmt.Errorf("get insert id error %w", err)
//	}
//
//	return id, nil
//}
//
//func (db *database) AuthUser(ctx context.Context, email string) (int64, string, error) {
//	row := db.client.QueryRowContext(ctx, authUser, email)
//
//	var id int64
//	var hash string
//
//	err := row.Scan(&id, &hash)
//
//	switch {
//	case err == sql.ErrNoRows:
//		return 0, "", UserNotFoundError
//	case err != nil:
//		return 0, "", err
//	}
//
//	return id, hash, nil
//}
