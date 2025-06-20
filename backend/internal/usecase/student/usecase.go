package usecase

import (
	"context"
	"golang_graphs/backend/internal/domain/model"
	repository "golang_graphs/backend/internal/domain/student/repository"
	studentservice "golang_graphs/backend/internal/domain/student/service"
	graphconverter "golang_graphs/backend/internal/domain/student/service/graphconverter"
	taskcheck "golang_graphs/backend/internal/domain/student/service/taskcheck"
	userservice "golang_graphs/backend/internal/domain/user/service"
)

type StudentUseCase interface {
	GetAssignedTasksByModule(ctx context.Context, in *GetAssignedTasksByModuleDTOIn) (*GetAssignedTasksByModuleDTOOut, error)
	AuthToken(ctx context.Context, token string) (*AuthTokenDTOOut, error)
	SendAnswers(ctx context.Context, request *SendAnswersDTOIn) (*SendAnswersDTOOut, error)
}

type studentUseCase struct {
	studentRepo    repository.StudentRepository
	studentService studentservice.StudentService
	taskChecker    taskcheck.Checker
	graphConverter graphconverter.GraphConverter
	userService    userservice.UserService
}

func NewStudentUseCase(
	repo repository.StudentRepository,
	studentService studentservice.StudentService,
	taskChecker taskcheck.Checker,
	graphConverter graphconverter.GraphConverter,
	userService userservice.UserService,
) StudentUseCase {
	return &studentUseCase{
		studentRepo:    repo,
		studentService: studentService,
		taskChecker:    taskChecker,
		graphConverter: graphConverter,
		userService:    userService,
	}
}

func (u *studentUseCase) GetAssignedTasksByModule(ctx context.Context, in *GetAssignedTasksByModuleDTOIn) (*GetAssignedTasksByModuleDTOOut, error) {
	user := &model.User{ID: in.UserID}
	module := &model.Module{ModuleId: in.ModuleID}
	out, err := u.studentRepo.GetAssignedTasksByModule(ctx, user, module)
	if err != nil {
		return nil, err
	}

	return &GetAssignedTasksByModuleDTOOut{Tasks: out}, nil
}

func (u *studentUseCase) AuthToken(ctx context.Context, token string) (*AuthTokenDTOOut, error) {
	user, err := u.userService.ParseToken(token)
	if err != nil {
		return nil, ErrParseToken
	}

	if user.Role != "student" {
		return nil, ErrNoPermissions
	}

	out, err := u.studentRepo.SelectStudent(ctx, user)
	if err != nil {
		return nil, err
	}

	return &AuthTokenDTOOut{
		UserID:    out.UserID,
		StudentID: out.ID,
	}, nil
}

func (u *studentUseCase) SendAnswers(ctx context.Context, request *SendAnswersDTOIn) (*SendAnswersDTOOut, error) {

	return &SendAnswersDTOOut{}, nil
}

func (u *studentUseCase) exportTaskCheck(ctx context.Context, taskType, subType string) {
	// var TaskChecks = map[string]any{
	// 	"LinearToLine": u.taskChecker.CheckLinearToLine,
	// }
}
