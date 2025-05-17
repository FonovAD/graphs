package config

import (
	storage "golang_graphs/backend/internal/infrastructure/storage/pg"
)

type Config struct {
	Host string `config:"APP_HOST" yaml:"host"`
	Port string `config:"APP_PORT" yaml:"port"`

	Postgres     storage.PGConfig `config:"postgres"`
	TestPostgres storage.PGConfig `config:"test_postgres"`
}
