package teacherrepository

import (
	"context"
	model "golang_graphs/backend/internal/domain/model/user"
)

type TeacherRepository interface {
	InsertUser(ctx context.Context, user *model.User) (*model.User, error)
}
