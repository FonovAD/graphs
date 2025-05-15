package userrepository

import (
	"context"
	model "golang_graphs/backend/internal/domain/model/user"
)

type UserRepository interface {
	SelectUserByEmail(ctx context.Context, email string) (*model.User, error)
}
