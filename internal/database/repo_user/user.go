package repo_user

import (
	"database/sql"
	"fmt"
	"golang_graphs/internal/database"
)

type user struct {
	client *sql.DB
}

type User interface {
	Get()
}

func NewRepoUser(config database.Config) (*user, error) {
	connInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database,
	)
	client, err := sql.Open("postgres", connInfo)

	if err != nil {
		return nil, err
	}

	return &user{client: client}, nil
}
