package usecase

import (
	"golang_graphs/backend/internal/domain/model"
	service "golang_graphs/backend/internal/domain/student/service/graphconverter"
	"time"
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
	TaskID int64 `json:"taskID"`
}

type SendAnswersDTOIn struct {
	UserID  int64
	LabID   int64    `json:"labID,omitempty"`
	Modules []Module `json:"modules"`
}

type Module struct {
	TypeID        int64                     `json:"type,omitempty"`
	TaskID        int64                     `json:"taskID,omitempty"`
	DataModule    []DataAnswer              `json:"data,omitempty"`
	RadiusAns     *int                      `json:"radiusAns,omitempty"`
	DiameterAns   *int                      `json:"diameterAns,omitempty"`
	Matrix1       map[string]map[string]int `json:"matrix1,omitempty"`
	Matrix2       map[string]map[string]int `json:"matrix2,omitempty"`
	Source        *string                   `json:"source,omitempty"`
	Target        *string                   `json:"target,omitempty"`
	WeightPathAns map[string]int            `json:"weightPathAns,omitempty"`
	MinPathAns    *int                      `json:"minPathAns,omitempty"`
	IsEulerAns    *bool                     `json:"isEulerAns,omitempty"`
	IsHamiltonian *bool                     `json:"isHamiltonianAns,omitempty"`
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
	LabID     int64     `json:"labID"`
	StartTime time.Time `json:"startTime"`
}

type FinishLabDTOIn struct {
	UserID int64
	LabID  int64 `json:"labID"`
}

type FinishLabDTOOut struct {
	LabID int64 `json:"labID"`
}
