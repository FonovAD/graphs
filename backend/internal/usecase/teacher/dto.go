package usecase

import (
	model "golang_graphs/backend/internal/domain/model"
	"time"

	"github.com/shopspring/decimal"
)

type CreateUserDTOIn struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	FatherName string `json:"fatherName"`
	Role       string `json:"role"`
}

type CreateUserDTOOut struct {
	UserID int64 `json:"userId"`
}

type GetModulesDTOOut struct {
	Modules []model.Module `json:"modules"`
}

type CreateLabDTOIn struct {
	LabName          string    `json:"name"`
	Description      string    `json:"description"`
	Duration         int64     `json:"duration"`
	RegistrationDate time.Time `json:"registrationDate"`
	TeacherID        int64     `json:"teacherId"`
}

type CreateLabDTOOut struct {
	LabID int64 `json:"labId"`
}

type GetLabInfoDTOIn struct {
	LabID int64 `json:"labId"`
}

type GetLabInfoDTOOut struct {
	LabName          string    `json:"name"`
	Description      string    `json:"description"`
	Duration         int64     `json:"duration"`
	RegistrationDate time.Time `json:"registrationDate"`
	TeacherID        int64     `json:"teacherId"`
	TeacherFIO       string    `json:"teacherFio"`
}

type RemoveUserLabDTOIn struct {
	UserID int64 `json:"userId"`
	LabID  int64 `json:"labId"`
}

type RemoveUserLabDTOOut struct {
	UserLabID int64 `json:"userLabId"`
}

type UpdateLabInfoDTOIn struct {
	LabID       int64  `json:"labId"`
	LabName     string `json:"labName"`
	Description string `json:"description"`
	Duration    int64  `json:"duration"`
}

type AssignLabDTOIn struct {
	UserID         int64     `json:"userId"`
	LabID          int64     `json:"labId"`
	AssignmentDate time.Time `json:"assignmentDate"`
	StartTime      time.Time `json:"startTime"`
	AssigneID      int64     `json:"assigneId"`
	Deadline       time.Time `json:"deadline"`
}

type AssignLabDTOOut struct {
	UserLabID int64 `json:"userLabId"`
}

type AssignLabGroupDTOIn struct {
	LabID          int64     `json:"labId"`
	AssignmentDate time.Time `json:"assignmentDate"`
	StartTime      time.Time `json:"startTime"`
	AssigneID      int64     `json:"assigneId"`
	Deadline       time.Time `json:"deadline"`
	GroupID        int64     `json:"groupId"`
}

type AssignLabGroupDTOOut struct {
	LabID int64 `json:"labId"`
}

type AddModuleLabDTOIn struct {
	LabID    int64           `json:"labId"`
	ModuleID int64           `json:"moduleId"`
	Weight   decimal.Decimal `json:"weight"`
}

type AddModuleLabDTOOut struct {
	ModuleLabID int64 `json:"moduleLabId"`
}

type RemoveModuleLabDTOIn struct {
	LabID    int64 `json:"labId"`
	ModuleID int64 `json:"moduleId"`
}

type RemoveModuleLabDTOOut struct {
	ModuleLabID int64 `json:"moduleLabId"`
}

type GetNonAssignedLabsDTOIn struct {
	Page int64 `json:"page"`
}

type GetNonAssignedLabsDTOOut struct {
	Labs []model.Lab `json:"labs"`
}

type GetAssignedLabsDTOIn struct {
	Page int64 `json:"page"`
}

type GetAssignedLabsDTOOut struct {
	Labs []model.UserLabWithInfo `json:"labs"`
}

type GetLabModulesDTOIn struct {
	LabID int64 `json:"labId"`
}

type GetLabModulesDTOOut struct {
	LabID   int64                `json:"labId"`
	Modules []model.ModulesInLab `json:"modules"`
}

type AuthTokenDTOOut struct {
	UserID    int64
	TeacherID int64
}

type GetGroupsDTOOut struct {
	Groups []model.Group `json:"groups"`
}
