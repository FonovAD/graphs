package repository

import (
	"context"
	"golang_graphs/backend/internal/domain/model"
)

type StudentRepository interface {
	GetAssignedTasksByModule(ctx context.Context, user *model.User, module *model.Module) ([]model.Task, error)
	SelectStudent(ctx context.Context, user *model.User) (*model.Student, error)
}
