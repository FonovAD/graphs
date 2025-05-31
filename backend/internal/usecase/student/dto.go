package usecase

import "golang_graphs/backend/internal/domain/model"

type GetAssignedTasksByModuleDTOIn struct {
	UserID   int64 `json:"userID"`
	ModuleID int64 `json:"moduleID"`
}

type GetAssignedTasksByModuleDTOOut struct {
	Tasks []model.Task `json:"tasks"`
}

type AuthTokenDTOOut struct {
	UserID    int64
	StudentID int64
}
