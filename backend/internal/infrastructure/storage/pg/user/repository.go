package pg

import (
	"context"
	"database/sql"
	"errors"
	model "golang_graphs/backend/internal/domain/model/user"
	userrepository "golang_graphs/backend/internal/domain/user/repository"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type userRepository struct {
	conn *sqlx.DB
}

func NewUserRepository(conn *sqlx.DB) userrepository.UserRepository {
	return &userRepository{conn}
}

func (r *userRepository) SelectUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	logrus.WithFields(logrus.Fields{"op": "infra.user.repo.SelectUserByEmail"}).Info("email: ", email)
	err := r.conn.QueryRowxContext(ctx, SelectUserByEmail, email).StructScan(&user)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}
