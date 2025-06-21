package usecase

import (
	"context"
	"fmt"
	"golang_graphs/backend/internal/domain/model"
	repository "golang_graphs/backend/internal/domain/student/repository"
	studentservice "golang_graphs/backend/internal/domain/student/service"
	graphconverter "golang_graphs/backend/internal/domain/student/service/graphconverter"
	taskcheck "golang_graphs/backend/internal/domain/student/service/taskcheck"
	userservice "golang_graphs/backend/internal/domain/user/service"
	"strings"
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
	userLab := &model.UserLab{UserID: request.UserID, LabID: request.LabID}
	score, err := u.studentRepo.SelectScore(ctx, userLab)
	if err != nil {
		return &SendAnswersDTOOut{}, err
	}

	if score.Score.Valid {
		return &SendAnswersDTOOut{}, fmt.Errorf("answer was already sent")
	}

	taskType, err := u.studentRepo.SelectModuleTypeByLab(ctx, userLab)
	if err != nil {
		return &SendAnswersDTOOut{}, err
	}

	taskType.TaskType = strings.TrimSpace(taskType.TaskType)

	return &SendAnswersDTOOut{}, nil
}

func (u *studentUseCase) extractTaskCheck(ctx context.Context, taskType string) any {
	// поход в бд
	var TaskChecks = map[string]any{
		"Реберный граф В":                      u.taskChecker.CheckLinearToLine,
		"Реберный граф Из":                     u.taskChecker.CheckLinearFromLine,
		"Радиус и диаметр":                     u.taskChecker.CheckRadiusAndDiameter,
		"Матрица смежности":                    u.taskChecker.CheckAdjacentMatrix,
		"Эйлеров граф":                         u.taskChecker.CheckEulerGraph,
		"Кратчайший путь":                      u.taskChecker.CheckMinPath,
		"Планарный граф":                       u.taskChecker.CheckPlanarGraph,
		"Бинарные операции Пересечение графов": u.taskChecker.CheckIntersectionGraphs,
		"Бинарные операции Объединение графов": u.taskChecker.CheckUnionGraphs,
		"Бинарные операции Соединение графов":  u.taskChecker.CheckJoinGraphs,
		"Харари":            u.taskChecker.CheckCartesianProduct,
		"Горбатов":          u.taskChecker.CheckTensorProduct,
		"Композиция графов": u.taskChecker.CheckLexicographicalProduct,
		"Бинарные операции Пересечение матриц": u.taskChecker.CheckIntersectionMatrices,
		"Бинарные операции Объединение матриц": u.taskChecker.CheckUnionMatrices,
		"Бинарные операции Соединение матриц":  u.taskChecker.CheckJoinMatrices,
		"Гамильтонов граф":                     u.taskChecker.CheckHamiltonian,
	}

	return TaskChecks[taskType]

}
