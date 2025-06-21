package repository

import (
	"context"
	"golang_graphs/backend/internal/domain/model"
)

type StudentRepository interface {
	GetAssignedTasksByModule(ctx context.Context, user *model.User, module *model.Module) ([]model.AssignedTaskByModule, error)
	SelectStudent(ctx context.Context, user *model.User) (*model.Student, error)
	SelectModuleTypeByTask(ctx context.Context, userTask *model.UserTask) (*model.TaskType, error)
	SelectModuleTypeByLab(ctx context.Context, userLab *model.UserLab) (*model.TaskType, error)
	SelectScore(ctx context.Context, userLab *model.UserLab) (*model.AssignedTaskByModule, error)
}
