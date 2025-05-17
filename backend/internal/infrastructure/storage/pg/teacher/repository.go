package storage

import (
	"context"
	"fmt"
	model "golang_graphs/backend/internal/domain/model/user"
	teacherrepository "golang_graphs/backend/internal/domain/teacher/repository"
	"golang_graphs/backend/internal/logger"

	"github.com/jmoiron/sqlx"
)

type teacherRepository struct {
	conn   *sqlx.DB
	logger logger.Logger
}

func NewTeacherRepository(conn *sqlx.DB, logger logger.Logger) teacherrepository.TeacherRepository {
	return &teacherRepository{conn: conn, logger: logger}
}

func (r *teacherRepository) InsertUser(ctx context.Context, user *model.User) (*model.User, error) {
	rows, err := r.conn.NamedQueryContext(ctx, insertIntoUsers, user)
	if err != nil {
		r.logger.LogDebug("infra.storage.pg.teacher.repo.InsertUser", err, user)

		return nil, fmt.Errorf("create user error %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		if err = rows.Scan(&user.Id); err != nil {
			r.logger.LogWarning("infra.storage.pg.teacher.repo.InsertUser", err, user)

			return nil, fmt.Errorf("scan userID error %w", err)
		}
	}

	return user, nil
}
