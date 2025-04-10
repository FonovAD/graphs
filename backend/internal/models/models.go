package models

import "golang_graphs/backend/internal/dto"

type AuthUserRequest struct {
	Email    string
	Password string
}

type AuthUserResponse struct {
	Token string
}

type CheckResultsRequest struct{}

type CheckResultsResponse struct {
	Results []dto.Result `json:"results"`
}

type CreateUserRequest struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
}

type CreateUserResponse struct {
	Token string
}

type GetTasksFromTestsRequest struct {
	TestID int64 `json:"test_id"`
}

type GetTasksFromTestsResponse struct {
	Tasks []dto.Task `json:"tasks"`
}

type GetTestsRequest struct{}

type GetTestsResponse struct {
	Tests []dto.Test `json:"tests"`
}

type SendAnswersRequest struct {
	Modules []Module `json:"modules"`
}

type SendAnswersResponse struct {
	Result dto.Result `json:"result"`
}

type SendTaskResultResponse struct {
	TaskType int64 `json:"result"`
}

type Module struct {
	TaskID     int64      `json:"type"`
	DataModule DataAnswer `json:"data"`
}

type DataAnswer struct {
	Nodes []NodeJSON `json:"nodes"`
	Edges []EdgeJSON `json:"edges"`
}

type BadRequestResponse struct {
	ErrorMsg string `json:"error_msg"`
}

type InternalServerErrorResponse struct {
	ErrorMsg string `json:"error_msg"`
}

type InsertTestRequest struct {
	Test dto.Test `json:"test"`
}

type InsertTestResponse struct {
	TestID int64 `json:"test_id"`
}

type InsertTaskRequest struct {
	Task dto.Task `json:"task"`
}

type InsertTaskResponse struct {
	TaskID int64 `json:"task_id"`
}
