package controller

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang_graphs/internal/database"
	"golang_graphs/internal/dto"
	"golang_graphs/internal/rest_models"
	"golang_graphs/pkg/create_random_string"
	"time"
)

type controller struct {
	db      database.Database
	creator create_random_string.Creator
}

func NewController(db database.Database, creator create_random_string.Creator) Controller {
	return &controller{db: db, creator: creator}
}

type Controller interface {
	CreateUser(ctx context.Context, user rest_models.CreateUserRequest) (dto.User, error)
	AuthUser(ctx context.Context, email, password string) (dto.User, error)
	GetTests(ctx context.Context) (rest_models.GetTestsResponse, error)
	GetTasksFromTest(ctx context.Context, testID int64) ([]dto.Task, error)
	CheckResults(ctx context.Context, userID int64) (rest_models.CheckResultsResponse, error)
	SendAnswers(ctx context.Context, answers []rest_models.Answer, testID int64) (dto.Result, error)
}

func (c *controller) SendAnswers(ctx context.Context, answers []rest_models.Answer, testID int64) (dto.Result, error) {
	tasksWithAnswers, err := c.GetTasksFromTest(ctx, testID)
	if err != nil {
		return dto.Result{}, nil
	}

	grade := int64(0)
	for _, answer := range answers {
		grade += c.checkResult(answer, c.findAnswerByID(tasksWithAnswers, answer.TaskID))
	}

	return dto.Result{
		Start:     time.Time{},
		End:       time.Now(),
		Grade:     grade,
		StudentID: 0,
		TestID:    testID,
	}, nil
}

func (c *controller) findAnswerByID(tasksWithAnswer []dto.Task, taskID int64) dto.Task {
	for _, task := range tasksWithAnswer {
		if task.ID == taskID {
			return task
		}
	}
	return dto.Task{}
}

func (c *controller) checkResult(answer rest_models.Answer, taskWithAnswer dto.Task) int64 {
	if answer.Answer == taskWithAnswer.Answer {
		return int64(taskWithAnswer.MaxGrade)
	}
	return 0
}

func (c *controller) CheckResults(ctx context.Context, userID int64) (rest_models.CheckResultsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	results, err := c.db.GetResultsByUserID(ctx, userID)
	if err != nil {
		return rest_models.CheckResultsResponse{}, err
	}

	return rest_models.CheckResultsResponse{Results: results}, nil
}

func (c *controller) CreateUser(ctx context.Context, user rest_models.CreateUserRequest) (dto.User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	salt := c.creator.RandomString()

	hash, err := hashPassword(user.Password, salt)
	if err != nil {
		return dto.User{}, err
	}

	userDto := dto.User{
		DateRegistration: time.Now(),
		Email:            user.Email,
		Password:         hash,
		FirstName:        user.FirstName,
		LastName:         user.LastName,
		Role:             "student",
		PasswordSalt:     salt,
	}

	userFromBD, err := c.db.InsertUser(ctx, userDto)
	if err != nil {
		return dto.User{}, err
	}

	return userFromBD, nil
}

func (c *controller) GetTests(ctx context.Context) (rest_models.GetTestsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	tests, err := c.db.GetTests(ctx)
	if err != nil {
		return rest_models.GetTestsResponse{}, err
	}

	return rest_models.GetTestsResponse{Tests: tests}, nil
}

func (c *controller) GetTasksFromTest(ctx context.Context, testID int64) ([]dto.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	tasks, err := c.db.GetTasksFromTest(ctx, testID)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (c *controller) AuthUser(ctx context.Context, email, password string) (dto.User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	user, err := c.db.SelectUserByEmail(ctx, email)
	if err != nil {
		return dto.User{}, err
	}

	password = password + user.PasswordSalt

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return dto.User{}, fmt.Errorf("incorrect nick or password %w", err)
	}

	return user, nil
}

// Hash password using the bcrypt hashing algorithm
func hashPassword(password, salt string) (string, error) {
	// Convert password string to byte slice
	var passwordBytes = []byte(password + salt)

	// Hash password with bcrypt's default cost
	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)

	return string(hashedPasswordBytes), err
}
