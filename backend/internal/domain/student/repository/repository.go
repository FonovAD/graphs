package repository

import (
	"context"
	"golang_graphs/backend/internal/domain/model"
)

type StudentRepository interface {
	GetAssignedTasksByModule(ctx context.Context, user *model.User, module *model.Module) ([]model.AssignedTaskByModule, error)
	SelectStudent(ctx context.Context, user *model.User) (*model.Student, error)
}
