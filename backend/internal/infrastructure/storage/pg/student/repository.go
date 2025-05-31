package storage

import (
	"context"
	"database/sql"
	"golang_graphs/backend/internal/domain/model"
	studentrepository "golang_graphs/backend/internal/domain/student/repository"
	"golang_graphs/backend/internal/logger"

	"github.com/jmoiron/sqlx"
)

type studentRepository struct {
	conn   *sqlx.DB
	logger logger.Logger
}

func NewStudentRepository(conn *sqlx.DB, logger logger.Logger) studentrepository.StudentRepository {
	return &studentRepository{conn: conn, logger: logger}
}

func (r *studentRepository) InsertTaskResult(ctx context.Context) (int64, error) {
	// 	log.Info("InsertTaskResult", result)
	// 	var id int64

	// 	row := r.conn.QueryRowxContext(ctx, InsertIntoTaskResult,
	// 		result.Type, result.UserID, result.Grade)

	// 	err := row.Scan(&id)

	// 	if err != nil {
	// 		if errors.Is(err, sql.ErrNoRows) {
	// 			return -1, fmt.Errorf("conflict on composite key (task_type, usersid)")
	// 		}
	// 		return 0, fmt.Errorf("insert task result error %w", err)
	// 	}

	// 	return id, nil
	// }
	return 0, nil
}

func (r *studentRepository) GetAssignedTasksByModule(ctx context.Context, user *model.User, module *model.Module) ([]model.TaskByModule, error) {
	var tasks []model.TaskByModule
	err := r.conn.SelectContext(ctx, &tasks, selectTasksByUserID, user.ID, module.ModuleId)
	if err != nil {
		r.logger.LogDebug(opSelectTasksByUserID, err, user.ID)
		return nil, err
	}

	return tasks, nil
}

func (r *studentRepository) SelectStudent(ctx context.Context, user *model.User) (*model.Student, error) {
	var student model.Student
	rows, err := r.conn.NamedQueryContext(ctx, selectStudent, user)
	if err != nil {
		r.logger.LogWarning(opSelectStudent, err, user.ID)
		return nil, err
	}

	if rows.Next() {
		if err := rows.Scan(&student.ID); err != nil {
			r.logger.LogWarning(opSelectStudent, err, user.ID)
			return nil, err
		}
		student.UserID = user.ID
		return &student, nil
	}
	r.logger.LogDebug(opSelectStudent, nil, user.ID)
	return nil, sql.ErrNoRows
}
