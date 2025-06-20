package usecase

import (
	"golang_graphs/backend/internal/domain/model"
	service "golang_graphs/backend/internal/domain/student/service/graphconverter"
)

type GetAssignedTasksByModuleDTOIn struct {
	UserID   int64 `json:"userID"`
	ModuleID int64 `json:"moduleID"`
}

type GetAssignedTasksByModuleDTOOut struct {
	Tasks []model.AssignedTaskByModule `json:"tasks"`
}

type AuthTokenDTOOut struct {
	UserID    int64
	StudentID int64
}

type SendAnswersDTOOut struct {
	TypeID int64
}

type SendAnswersDTOIn struct {
	Modules []Module `json:"modules"`
}

type Module struct {
	TypeID      int64      `json:"type"`
	SubType     string     `json:"subType"`
	DataModule  DataAnswer `json:"data"`
	InputValue1 string     `json:"inputValue1"`
	InputValue2 string     `json:"inputValue2"`
	InputValue3 string     `json:"inputValue3"`
	InputValue4 string     `json:"inputValue4"`
	InputValue5 string     `json:"inputValue5"`
	InputValue6 string     `json:"inputValue6"`
}

type DataAnswer struct {
	Nodes []service.NodeJSON `json:"nodes"`
	Edges []service.EdgeJSON `json:"edges"`
}
