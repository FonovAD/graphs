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
	UserID  int64
	LabID   int64    `json:"labID"`
	Modules []Module `json:"modules"`
}

type Module struct {
	TypeID        int64                     `json:"type"`
	TaskID        int64                     `json:"taskID"`
	DataModule    []DataAnswer              `json:"data"`
	RadiusAns     *int                      `json:"radiusAns"`
	DiameterAns   *int                      `json:"diameterAns"`
	Matrix1       map[string]map[string]int `json:"matrix1"`
	Matrix2       map[string]map[string]int `json:"matrix2"`
	Source        *string                   `json:"source"`
	Target        *string                   `json:"target"`
	WeightPathAns map[string]int            `json:"weightPathAns"`
	MinPathAns    *int                      `json:"minPathAns"`
	IsEulerAns    *bool                     `json:"isEulerAns"`
	IsHamiltonian *bool                     `json:"isHamiltonianAns"`
}

type DataAnswer struct {
	Nodes []service.NodeJSON `json:"nodes"`
	Edges []service.EdgeJSON `json:"edges"`
}

type BeginLabDTOIn struct {
	UserID int64
	LabID  int64 `json:"labID"`
}

type BeginLabDTOOut struct {
	LabID int64 `json:"labID"`
}

type FinishLabDTOIn struct {
	UserID int64
	LabID  int64 `json:"labID"`
}

type FinishLabDTOOut struct {
	LabID int64 `json:"labID"`
}
