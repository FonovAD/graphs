package pg

import (
	"context"
	studentrepository "golang_graphs/backend/internal/domain/student/repository"

	"github.com/jmoiron/sqlx"
)

type studentRepository struct {
	conn *sqlx.DB
}

func NewStudentRepository(conn *sqlx.DB) studentrepository.StudentRepository {
	return &studentRepository{conn}
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
