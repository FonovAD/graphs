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

// func (c *controller) SendAnswers(ctx context.Context, user dto.User, request models.SendAnswersRequest) (models.SendTaskResultResponse, error) {
// 	grade := int64(0)
// 	moduleType := int64(0)
// 	for _, module := range request.Modules {
// 		if len(module.DataModule.Nodes) > 0 && len(module.DataModule.Edges) > 0 {
// 			grade = 100
// 		}
// 		moduleType = module.TaskID
// 		// grade += c.checkResult(answer, c.findAnswerByID(tasksWithAnswers, answer.TaskID))
// 	}

// 	// maxGrade := int64(100)
// 	// for _, answer := range tasksWithAnswers {
// 	// 	maxGrade += answer.MaxGrade
// 	// }

// 	// result := dto.Result{
// 	// 	Start:     time.Time{},
// 	// 	End:       time.Now(),
// 	// 	Grade:     grade,
// 	// 	StudentID: user.Id,
// 	// 	TestID:    1,
// 	// 	MaxGrade:  maxGrade,
// 	// }

// 	// err := c.db.InsertResult(ctx, result)
// 	result := dto.TaskResult{
// 		Type:   moduleType,
// 		UserID: user.Id,
// 		Grade:  grade,
// 	}

// 	id, err := c.db.InsertTaskResult(ctx, result)
// 	if err != nil && id == -1 {
// 		return models.SendTaskResultResponse{TaskType: id}, err
// 	}

// 	if err != nil {
// 		return models.SendTaskResultResponse{}, err
// 	}

// 	return models.SendTaskResultResponse{TaskType: id}, nil
// }
