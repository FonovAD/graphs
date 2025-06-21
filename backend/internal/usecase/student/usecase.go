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
	"time"
)

type StudentUseCase interface {
	GetAssignedTasksByModule(ctx context.Context, in *GetAssignedTasksByModuleDTOIn) (*GetAssignedTasksByModuleDTOOut, error)
	AuthToken(ctx context.Context, token string) (*AuthTokenDTOOut, error)
	SendAnswers(ctx context.Context, request *SendAnswersDTOIn) (*SendAnswersDTOOut, error)
	BeginLab(ctx context.Context, in *BeginLabDTOIn) (*BeginLabDTOOut, error)
	FinishLab(ctx context.Context, in *FinishLabDTOIn) (*FinishLabDTOOut, error)
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

func (u *studentUseCase) SendAnswers(ctx context.Context, in *SendAnswersDTOIn) (*SendAnswersDTOOut, error) {
	userLab := &model.UserLab{UserID: in.UserID, LabID: in.LabID}
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
	module := in.Modules[0]
	taskType.TaskType = strings.TrimSpace(taskType.TaskType)
	inputData := &taskcheck.InputData{
		AnswerGraph:      u.safeGetGraph(ctx, module.DataModule, 0),
		TaskGraph1:       u.safeGetGraph(ctx, module.DataModule, 1),
		TaskGraph2:       u.safeGetGraph(ctx, module.DataModule, 2),
		RadiusAns:        safeGetPointerInt(module.RadiusAns),
		DiameterAns:      safeGetPointerInt(module.DiameterAns),
		Matrix1:          module.Matrix1,
		Matrix2:          module.Matrix2,
		Source:           safeGetPointerString(module.Source),
		Target:           safeGetPointerString(module.Target),
		WeightsPathAns:   module.WeightPathAns,
		MinPathAns:       safeGetPointerInt(module.MinPathAns),
		IsEulerAns:       safeGetPointerBool(module.IsEulerAns),
		IsHamiltonianAns: safeGetPointerBool(module.IsHamiltonian),
	}

	taskCheckerFunc := u.extractTaskCheck(taskType.TaskType)

	targetScore := taskCheckerFunc(inputData)
	return &SendAnswersDTOOut{TypeID: int64(targetScore)}, nil
}

func (u *studentUseCase) BeginLab(ctx context.Context, in *BeginLabDTOIn) (*BeginLabDTOOut, error) {
	userLab := &model.UserLab{UserID: in.UserID, LabID: in.LabID, StartTime: time.Now()}
	out, err := u.studentRepo.BeginLab(ctx, userLab)
	if err != nil {
		return &BeginLabDTOOut{}, err
	}

	return &BeginLabDTOOut{LabID: out.LabID}, nil
}

func (u *studentUseCase) FinishLab(ctx context.Context, in *FinishLabDTOIn) (*FinishLabDTOOut, error) {
	userLab := &model.UserLab{UserID: in.UserID, LabID: in.LabID}
	out, err := u.studentRepo.FinishLab(ctx, userLab)
	if err != nil {
		return &FinishLabDTOOut{}, err
	}

	return &FinishLabDTOOut{LabID: out.LabID}, nil
}

func (u *studentUseCase) extractTaskCheck(taskType string) func(*taskcheck.InputData) int {
	var TaskChecks = map[string]func(*taskcheck.InputData) int{
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

func (u *studentUseCase) safeGetGraph(ctx context.Context, data []DataAnswer, index int) *model.Graph {
	if index < 0 || index >= len(data) {
		return nil
	}
	graph, err := u.graphConverter.ConvertJSONStructsToGraph(ctx, data[index].Nodes, data[index].Edges)
	if err != nil {
		return nil
	}

	return graph
}

func safeGetPointerInt(ptr *int) int {
	if ptr == nil {
		return 0
	}

	return *ptr
}

func safeGetPointerString(ptr *string) string {
	if ptr == nil {
		return ""
	}

	return *ptr
}

func safeGetPointerBool(ptr *bool) bool {
	if ptr == nil {
		return false
	}

	return *ptr
}
