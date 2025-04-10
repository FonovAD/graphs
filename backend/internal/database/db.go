package database

import (
	"database/sql"
	"fmt"
	"golang_graphs/backend/internal/dto"
	"sync"

	_ "github.com/lib/pq"
)

var (
	UserNotFoundError = fmt.Errorf("user doesn't exist")
)

type database struct {
	client        *sql.DB
	cacheGetTests []dto.Test
	mu            *sync.Mutex
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

	return &database{client: client, cacheGetTests: nil, mu: &sync.Mutex{}}, nil
}
