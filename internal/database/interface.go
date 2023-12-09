package database

import (
	"context"
	"golang_graphs/internal/dto"
	"time"
)

type Database interface {
	Ping(ctx context.Context) error
	// добавляет к модели user поле ID
	InsertUser(ctx context.Context, user dto.User) (dto.User, error)
	SelectUserByEmail(ctx context.Context, email string) (dto.User, error)
	GetTests(ctx context.Context) ([]dto.Test, error)
	GetTasksFromTest(ctx context.Context, testID int64) ([]dto.Task, error)
	GetResultsByUserID(ctx context.Context, userID int64) ([]dto.Result, error)
	InsertResult(ctx context.Context, result dto.Result) error
}

func (db *database) Ping(ctx context.Context) error {
	return nil
}

func (db *database) InsertUser(ctx context.Context, user dto.User) (dto.User, error) {
	return dto.User{
		Id:               1,
		DateRegistration: time.Time{},
		Email:            "123213",
		Password:         "33",
		FirstName:        "2",
		LastName:         "3123",
		Role:             "asada",
		PasswordSalt:     "zxczxcasf",
	}, nil
}

func (db *database) SelectUserByEmail(ctx context.Context, email string) (dto.User, error) {
	return dto.User{
		Id:               1,
		DateRegistration: time.Time{},
		Email:            "123213",
		Password:         "33",
		FirstName:        "2",
		LastName:         "3123",
		Role:             "asada",
		PasswordSalt:     "zxczxcasf",
	}, nil
}

func (db *database) GetTests(ctx context.Context) ([]dto.Test, error) {
	return []dto.Test{{
		ID:       228,
		Name:     "228",
		Start:    time.Time{},
		End:      time.Time{},
		Interval: time.Time{},
	}}, nil
}
func (db *database) GetTasksFromTest(ctx context.Context, testID int64) ([]dto.Task, error) {
	return []dto.Task{{
		ID:       228,
		Name:     "228",
		Data:     "",
		Answer:   "",
		MaxGrade: 1,
	}}, nil
}
func (db *database) GetResultsByUserID(ctx context.Context, userID int64) ([]dto.Result, error) {
	return []dto.Result{{
		ID:        228,
		Start:     time.Time{},
		End:       time.Time{},
		Grade:     228,
		StudentID: 1,
		TestID:    1,
	}}, nil
}
func (db *database) InsertResult(ctx context.Context, result dto.Result) error {
	return nil
}
