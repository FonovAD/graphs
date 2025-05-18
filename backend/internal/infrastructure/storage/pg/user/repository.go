package storage

import (
	"context"
	"database/sql"
	"errors"
	model "golang_graphs/backend/internal/domain/model"
	userrepository "golang_graphs/backend/internal/domain/user/repository"
	"golang_graphs/backend/internal/logger"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	conn   *sqlx.DB
	logger logger.Logger
}

func NewUserRepository(conn *sqlx.DB, logger logger.Logger) userrepository.UserRepository {
	return &userRepository{conn: conn, logger: logger}
}

func (r *userRepository) SelectUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.conn.QueryRowxContext(ctx, SelectUserByEmail, email).StructScan(&user)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.logger.LogInfo("infra.user.repo.SelectUserByEmail", err, email)
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}
