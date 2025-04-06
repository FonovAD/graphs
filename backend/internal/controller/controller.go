package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"golang_graphs/backend/internal/controller/task_check"
	"golang_graphs/backend/internal/database"
	"golang_graphs/backend/internal/dto"
	"golang_graphs/backend/internal/models"
	"golang_graphs/backend/pkg/auth"
	"golang_graphs/backend/pkg/create_random_string"
	"log"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type controller struct {
	db          database.Database
	creator     create_random_string.Creator
	authService auth.Service
	ansChecker  task_check.Checker
}

func NewController(db database.Database, creator create_random_string.Creator, authService auth.Service, ansChecker task_check.Checker) Controller {
	return &controller{db: db, creator: creator, authService: authService, ansChecker: ansChecker}
}

type Controller interface {
	CreateUser(ctx context.Context, user models.CreateUserRequest) (models.CreateUserResponse, error)
	AuthUser(ctx context.Context, request models.AuthUserRequest) (models.AuthUserResponse, error)
	GetTests(ctx context.Context) (models.GetTestsResponse, error)
	GetTasksFromTest(ctx context.Context, request models.GetTasksFromTestsRequest) (models.GetTasksFromTestsResponse, error)
	CheckResults(ctx context.Context, user dto.User, request models.CheckResultsRequest) (models.CheckResultsResponse, error)
	SendAnswers(ctx context.Context, user dto.User, request models.SendAnswersRequest) (models.SendAnswersResponse, error)
	InsertTest(ctx context.Context, request models.InsertTestRequest) (models.InsertTestResponse, error)
	InsertTask(ctx context.Context, request models.InsertTaskRequest) (models.InsertTaskResponse, error)
	AuthToken(ctx context.Context, token string) (dto.User, error)
}

func (c *controller) AuthToken(ctx context.Context, token string) (dto.User, error) {
	user, err := c.authService.ParseToken(token)
	if err != nil {
		return dto.User{}, err
	}

	return user, nil
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

func (c *controller) parserNodeStruct(ctx context.Context, json_str string) (*models.NodesJSON, error) {
	ans := new(models.NodesJSON)
	var parser_ans []models.NodeJSON
	if err := json.Unmarshal([]byte(json_str), &parser_ans); err != nil {
		return nil, err
	}
	ans.NodeArr = parser_ans
	return ans, nil
}

func (c *controller) parserEdgeStruct(ctx context.Context, json_str string) (*models.EdgesJSON, error) {
	ans := new(models.EdgesJSON)
	var parser_ans []models.EdgeJSON
	if err := json.Unmarshal([]byte(json_str), &parser_ans); err != nil {
		return nil, err
	}
	ans.EdgeArr = parser_ans
	return ans, nil
}

func (c *controller) convertJSONStructsToGraph(ctx context.Context, nodes_json *models.NodesJSON, edges_json *models.EdgesJSON) (*models.Graph, error) {
	graph := new(models.Graph)
	node_id_map := make(map[string]int)
	curr_id := 0
	for _, node_json := range nodes_json.NodeArr {
		node_id_map[node_json.NodeData.Id] = curr_id
		curr_id++
		weight, err := strconv.Atoi(node_json.NodeData.Weight)
		if err != nil {
			weight = 0
		}
		graph.AddNodeByInfo(curr_id,
			node_json.NodeData.Label,
			node_json.NodeData.Color,
			weight,
			node_json.Position.X,
			node_json.Position.Y)
	}

	for _, edge := range edges_json.EdgeArr {
		src, err := graph.FindNodeById(node_id_map[edge.EdgeData.Source])
		if err != nil {
			return graph, err
		}
		trg, _ := graph.FindNodeById(node_id_map[edge.EdgeData.Target])
		if err != nil {
			return graph, err
		}
		graph.AddEdgeByInfo(
			src,
			trg,
			edge.EdgeData.Id,
			edge.EdgeData.Label,
			edge.EdgeData.Color,
			0)
	}
	return graph, nil
}

// func (c *controller) parserGraphStruct(ctx context.Context, request models.SendAnswersRequest) (*models.Graph, error) {
// 	graph := new(models.Graph)
// 	var json_struct
// 	// пробегаться по циклу request.Answers?
// 	if err := json.Unmarshal([]byte(request.Answers[0].Answer), &json_struct); err != nil {
// 		return nil, err
// 	}
// 	for _, data := range json_struct {
// 		if data.source == nil {

// 		}
// 	}
// }
// нужны парсеры для разных заданий
// узнать структуру всего json, приходящего с фронта, чтобы вытащить всю инфу по заданию
// в зависимости от задания (модуля) вызывать нужную функцию

func (c *controller) SendAnswers(ctx context.Context, user dto.User, request models.SendAnswersRequest) (models.SendAnswersResponse, error) {

	grade := int64(0)
	for _, module := range request.Modules {
		if len(module.DataModule.Nodes) > 0 && len(module.DataModule.Edges) > 0 {
			grade = 100
		}
		// grade += c.checkResult(answer, c.findAnswerByID(tasksWithAnswers, answer.TaskID))
	}

	maxGrade := int64(100)
	// for _, answer := range tasksWithAnswers {
	// 	maxGrade += answer.MaxGrade
	// }

	result := dto.Result{
		Start:     time.Time{},
		End:       time.Now(),
		Grade:     grade,
		StudentID: user.Id,
		TestID:    1,
		MaxGrade:  maxGrade,
	}

	err := c.db.InsertResult(ctx, result)
	if err != nil {
		return models.SendAnswersResponse{}, err
	}

	return models.SendAnswersResponse{Result: result}, nil
}

// func (c *controller) checkResult(answer models.Answer, taskWithAnswer dto.Task) int64 {
// 	if answer.Answer == taskWithAnswer.Answer {
// 		return int64(taskWithAnswer.MaxGrade)
// 	}
// 	return 0
// }

func (c *controller) findAnswerByID(tasksWithAnswer []dto.Task, taskID int64) dto.Task {
	for _, task := range tasksWithAnswer {
		if task.ID == taskID {
			return task
		}
	}
	return dto.Task{}
}

func (c *controller) CheckResults(ctx context.Context, user dto.User, request models.CheckResultsRequest) (models.CheckResultsResponse, error) {
	log.Printf("Check Result request %s", request)

	results, err := c.db.GetResultsByUserID(ctx, user.Id)
	if err != nil {
		return models.CheckResultsResponse{}, errors.Wrap(err, "error getting results from DB")
	}

	return models.CheckResultsResponse{Results: results}, nil
}

func (c *controller) CreateUser(ctx context.Context, user models.CreateUserRequest) (models.CreateUserResponse, error) {
	if err := ValidateCreateUser(user); err != nil {
		return models.CreateUserResponse{}, err
	}

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
	tests, err := c.db.GetTests(ctx)
	if err != nil {
		return models.GetTestsResponse{}, err
	}

	return models.GetTestsResponse{Tests: tests}, nil
}

func (c *controller) GetTasksFromTest(ctx context.Context, request models.GetTasksFromTestsRequest) (models.GetTasksFromTestsResponse, error) {
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
	if err := ValidateAuthUser(request); err != nil {
		return models.AuthUserResponse{}, err
	}
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
