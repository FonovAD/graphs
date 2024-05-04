package controller

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"golang_graphs/internal/database"
	"golang_graphs/internal/dto"
	"golang_graphs/internal/models"
	"golang_graphs/pkg/auth"
	"golang_graphs/pkg/create_random_string"
	"log"
	"time"
)

type controller struct {
	db          database.Database
	creator     create_random_string.Creator
	authService auth.Service
}

func NewController(db database.Database, creator create_random_string.Creator, authService auth.Service) Controller {
	return &controller{db: db, creator: creator, authService: authService}
}

type Controller interface {
	CreateUser(ctx context.Context, user models.CreateUserRequest) (models.CreateUserResponse, error)
	AuthUser(ctx context.Context, request models.AuthUserRequest) (models.AuthUserResponse, error)
	GetTests(ctx context.Context) (models.GetTestsResponse, error)
	GetTasksFromTest(ctx context.Context, request models.GetTasksFromTestsRequest) (models.GetTasksFromTestsResponse, error)
	CheckResults(ctx context.Context, request models.CheckResultsRequest) (models.CheckResultsResponse, error)
	SendAnswers(ctx context.Context, request models.SendAnswersRequest) (models.SendAnswersResponse, error)
	InsertTest(ctx context.Context, request models.InsertTestRequest) (models.InsertTestResponse, error)
	InsertTask(ctx context.Context, request models.InsertTaskRequest) (models.InsertTaskResponse, error)
}

func (c *controller) InsertTest(ctx context.Context, request models.InsertTestRequest) (models.InsertTestResponse, error) {
	id, err := c.db.InsertTest(ctx, request.Test)
	if err != nil {
		return models.InsertTestResponse{}, errors.Wrap(err, "error inserting test into database")
	}
	return models.InsertTestResponse{TestID: id}, nil
}

func (c *controller) InsertTask(ctx context.Context, request models.InsertTaskRequest) (models.InsertTaskResponse, error) {
	id, err := c.db.InsertTask(ctx, request.Task)
	if err != nil {
		return models.InsertTaskResponse{}, errors.Wrap(err, "error inserting task into database")
	}
	return models.InsertTaskResponse{TaskID: id}, nil
}

func (c *controller) SendAnswers(ctx context.Context, request models.SendAnswersRequest) (models.SendAnswersResponse, error) {
	tasksWithAnswers, err := c.db.GetTasksFromTest(ctx, request.TestID)
	if err != nil {
		return models.SendAnswersResponse{}, errors.Wrap(err, "error getting tasks from test ID")
	}

	grade := int64(0)
	for _, answer := range request.Answers {
		grade += c.checkResult(answer, c.findAnswerByID(tasksWithAnswers, answer.TaskID))
	}

	maxGrade := int64(0)
	for _, answer := range tasksWithAnswers {
		maxGrade += answer.MaxGrade
	}

	user, err := c.authService.AuthUser(request.Token)
	if err != nil {
		return models.SendAnswersResponse{}, errors.Wrap(err, "error getting user")
	}

	result := dto.Result{
		Start:     time.Time{},
		End:       time.Now(),
		Grade:     grade,
		StudentID: user.Id,
		TestID:    request.TestID,
		MaxGrade:  maxGrade,
	}

	err = c.db.InsertResult(ctx, result)
	if err != nil {
		return models.SendAnswersResponse{}, err
	}

	return models.SendAnswersResponse{Result: result}, nil
}

func (c *controller) checkResult(answer models.Answer, taskWithAnswer dto.Task) int64 {
	if answer.Answer == taskWithAnswer.Answer {
		return int64(taskWithAnswer.MaxGrade)
	}
	return 0
}

func (c *controller) findAnswerByID(tasksWithAnswer []dto.Task, taskID int64) dto.Task {
	for _, task := range tasksWithAnswer {
		if task.ID == taskID {
			return task
		}
	}
	return dto.Task{}
}

func (c *controller) CheckResults(ctx context.Context, request models.CheckResultsRequest) (models.CheckResultsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	log.Printf("Check Result request %s", request)

	user, err := c.authService.AuthUser(request.Token)
	if err != nil {
		return models.CheckResultsResponse{}, errors.Wrap(err, "error getting user")
	}

	results, err := c.db.GetResultsByUserID(ctx, user.Id)
	if err != nil {
		return models.CheckResultsResponse{}, errors.Wrap(err, "error getting results from DB")
	}

	return models.CheckResultsResponse{Results: results}, nil
}

func (c *controller) CreateUser(ctx context.Context, user models.CreateUserRequest) (models.CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	salt := c.creator.RandomString()

	hash, err := hashPassword(user.Password, salt)
	if err != nil {
		return models.CreateUserResponse{}, errors.Wrap(err, "hash password")
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

	userFromDB, err := c.db.InsertUser(ctx, userDto)
	if err != nil {
		return models.CreateUserResponse{}, errors.Wrap(err, "insert user")
	}

	token, err := c.authService.CreateToken(userFromDB)
	if err != nil {
		return models.CreateUserResponse{}, errors.Wrap(err, "error creating token")
	}

	return models.CreateUserResponse{Token: token}, nil
}

func (c *controller) GetTests(ctx context.Context) (models.GetTestsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	tests, err := c.db.GetTests(ctx)
	if err != nil {
		return models.GetTestsResponse{}, err
	}

	return models.GetTestsResponse{Tests: tests}, nil
}

// TODO
func (c *controller) GetTasksFromTest(ctx context.Context, request models.GetTasksFromTestsRequest) (models.GetTasksFromTestsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	tasks, err := c.db.GetTasksFromTest(ctx, request.TestID)
	if err != nil {
		return models.GetTasksFromTestsResponse{}, err
	}

	// удаляю ответы
	for i := 0; i < len(tasks); i++ {
		tasks[i].Answer = ""
	}

	return models.GetTasksFromTestsResponse{Tasks: tasks}, nil
}

func (c *controller) AuthUser(ctx context.Context, request models.AuthUserRequest) (models.AuthUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	user, err := c.db.SelectUserByEmail(ctx, request.Email)
	if err != nil {
		return models.AuthUserResponse{}, err
	}

	password := request.Password + user.PasswordSalt

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return models.AuthUserResponse{}, fmt.Errorf("incorrect nick or password %w", err)
	}

	token, err := c.authService.CreateToken(user)
	if err != nil {
		return models.AuthUserResponse{}, errors.Wrap(err, "error creating token")
	}

	return models.AuthUserResponse{Token: token}, nil
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
