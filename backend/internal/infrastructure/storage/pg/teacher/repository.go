package pg

import (
	"context"
	"fmt"
	model "golang_graphs/backend/internal/domain/model/user"
	teacherrepository "golang_graphs/backend/internal/domain/teacher/repository"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type teacherRepository struct {
	conn *sqlx.DB
}

func NewTeacherRepository(conn *sqlx.DB) teacherrepository.TeacherRepository {
	return &teacherRepository{conn}
}

func (r *teacherRepository) InsertUser(ctx context.Context, user *model.User) (*model.User, error) {
	log.Info("InsertUser", user)
	var id int64

	row, err := r.conn.NamedQueryContext(ctx, insertIntoUsers, &user)
	if err != nil {
		return nil, fmt.Errorf("create user error %w", err)
	}
	defer row.Close()

	err = row.Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("create user error %w", err)
	}

	user.Id = id

	return user, nil
}
