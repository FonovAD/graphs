package database

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
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
	log.Info("InsertUser", user)
	var id int64

	row := db.client.QueryRowContext(ctx, insertIntoUsers,
		user.Role, user.FirstName, user.LastName, user.Email, "", user.Password, user.PasswordSalt, user.DateRegistration)

	err := row.Scan(&id)

	if err != nil {
		return dto.User{}, fmt.Errorf("create user error %w", err)
	}

	user.Id = id

	return user, nil
}

func (db *database) SelectUserByEmail(ctx context.Context, email string) (dto.User, error) {
	log.Info("SelectUserByEmail", email)
	row := db.client.QueryRowContext(ctx, SelectUserByEmail, email)

	user := dto.User{
		Email: email,
	}

	fatherName := ""

	err := row.Scan(&user.Role, &user.FirstName, &user.LastName, &user.Email, &fatherName, &user.Password, &user.PasswordSalt, &user.DateRegistration)
	if err != nil {
		return dto.User{}, fmt.Errorf("select user by email error %w", err)
	}

	return user, nil
}

func (db *database) GetTests(ctx context.Context) ([]dto.Test, error) {
	log.Info("GetTests")
	rows, err := db.client.QueryContext(ctx, SelectAllAvailableTests)
	if err != nil {
		return nil, fmt.Errorf("select error %w", err)
	}
	defer rows.Close()

	tests := make([]dto.Test, 0)

	for rows.Next() {
		test := dto.Test{}
		if err := rows.Scan(&test.ID, &test.Name, &test.Start, &test.End); err != nil {
			return nil, fmt.Errorf("get tests row scan error %w", err)
		}
		tests = append(tests, test)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("get tests rows err %w", err)
	}

	return tests, nil
}

func (db *database) GetTasksFromTest(ctx context.Context, testID int64) ([]dto.Task, error) {
	log.Info("GetTasksFromTest", testID)
	return []dto.Task{{
		ID:       228,
		Name:     "228",
		Data:     "",
		Answer:   "",
		MaxGrade: 1,
	}}, nil
}

func (db *database) GetResultsByUserID(ctx context.Context, userID int64) ([]dto.Result, error) {
	log.Info("GetResultByUserID", userID)
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
	log.Info("InsertResult", result)
	var id int64

	row := db.client.QueryRowContext(ctx, InsertIntoResult,
		result.Start, result.End, result.Grade, result.StudentID, result.TestID)

	err := row.Scan(&id)

	if err != nil {
		return fmt.Errorf("insert result error %w", err)
	}

	return nil
}
