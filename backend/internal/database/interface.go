package database

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"golang_graphs/backend/internal/dto"
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
	InsertTest(ctx context.Context, test dto.Test) (int64, error)
	InsertTask(ctx context.Context, task dto.Task) (int64, error)
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
	//db.mu.Lock()
	//defer db.mu.Unlock()
	//if db.cacheGetTests != nil {
	//	return db.cacheGetTests, nil
	//}
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

	//db.cacheGetTests = tests

	return tests, nil
}

func (db *database) GetTasksFromTest(ctx context.Context, testID int64) ([]dto.Task, error) {
	rows, err := db.client.QueryContext(ctx, SelectTasksByTestID, testID)
	if err != nil {
		return nil, fmt.Errorf("select error %w", err)
	}
	defer rows.Close()

	tasks := make([]dto.Task, 0)

	for rows.Next() {
		task := dto.Task{}
		if err := rows.Scan(&task.ID, &task.Name, &task.Answer, &task.Data, &task.MaxGrade, &task.Description); err != nil {
			return nil, fmt.Errorf("get tasks row scan error %w", err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("get tasks rows err %w", err)
	}

	return tasks, nil
}

func (db *database) GetResultsByUserID(ctx context.Context, userID int64) ([]dto.Result, error) {
	rows, err := db.client.QueryContext(ctx, SelectResultsByUserID, userID)
	if err != nil {
		return nil, fmt.Errorf("select error %w", err)
	}
	defer rows.Close()

	results := make([]dto.Result, 0)

	for rows.Next() {
		result := dto.Result{}
		if err := rows.Scan(&result.ID, &result.Start, &result.End, &result.Grade, &result.MaxGrade, &result.StudentID, &result.TestID); err != nil {
			return nil, fmt.Errorf("get results row scan error %w", err)
		}
		results = append(results, result)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("get results rows err %w", err)
	}

	return results, nil
}

func (db *database) InsertResult(ctx context.Context, result dto.Result) error {
	log.Info("InsertResult", result)
	var id int64

	row := db.client.QueryRowContext(ctx, InsertIntoResult,
		result.Start, result.End, result.Grade, result.MaxGrade, result.StudentID, result.TestID)

	err := row.Scan(&id)

	if err != nil {
		return fmt.Errorf("insert result error %w", err)
	}

	return nil
}

func (db *database) InsertTest(ctx context.Context, test dto.Test) (int64, error) {
	//db.mu.Lock()
	//defer db.mu.Unlock()
	//db.cacheGetTests = nil

	log.Info("InsertTest", test)
	var id int64

	row := db.client.QueryRowContext(ctx, InsertTest, test.Name, test.Description, "24h", test.Start, test.End)
	err := row.Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("insert result error %w", err)
	}

	return id, nil
}

func (db *database) InsertTask(ctx context.Context, task dto.Task) (int64, error) {
	log.Info("InsertTask", task)
	var id int64

	row := db.client.QueryRowContext(ctx, InsertTask, task.TestID, task.Name, task.Answer, task.Data, task.MaxGrade, task.Description)
	err := row.Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("insert result error %w", err)
	}

	return id, nil
}
