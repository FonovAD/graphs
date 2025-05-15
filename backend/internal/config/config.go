package config

import (
	"golang_graphs/backend/internal/infrastructure/storage/pg"
)

type Config struct {
	Host string `config:"APP_HOST" yaml:"host"`
	Port string `config:"APP_PORT" yaml:"port"`

	Postgres     pg.DBConfig `config:"postgres"`
	TestPostgres pg.DBConfig `config:"test_postgres"`
}
