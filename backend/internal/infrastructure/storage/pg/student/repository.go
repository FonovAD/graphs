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

func (r *studentRepository) GetAssignedTasksByModule(ctx context.Context, user *model.User, module *model.Module) ([]model.AssignedTaskByModule, error) {
	var tasks []model.AssignedTaskByModule
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

func (r *studentRepository) SelectModuleTypeByLab(ctx context.Context, userLab *model.UserLab) (*model.TaskType, error) {
	var taskType model.TaskType
	rows, err := r.conn.NamedQueryContext(ctx, selectModuleTypeByLab, userLab)
	if err != nil {
		r.logger.LogWarning(opSelectModuleTypeByLab, err, map[string]any{"userID": userLab.UserID, "labID": userLab.LabID})
		return nil, err
	}

	if rows.Next() {
		if err := rows.Scan(&taskType.TaskType); err != nil {
			r.logger.LogWarning(opSelectModuleTypeByLab, err, map[string]any{"userID": userLab.UserID, "labID": userLab.LabID})
			return nil, err
		}
		return &taskType, nil
	}
	return nil, sql.ErrNoRows
}

func (r *studentRepository) SelectModuleTypeByTask(ctx context.Context, userTask *model.UserTask) (*model.TaskType, error) {
	var taskType model.TaskType
	rows, err := r.conn.NamedQueryContext(ctx, selectModuleTypeByTask, userTask)
	if err != nil {
		r.logger.LogWarning(opSelectModuleTypeByTask, err, map[string]any{"userID": userTask.UserID, "labID": userTask.TaskID})
		return nil, err
	}

	if rows.Next() {
		if err := rows.Scan(&taskType.TaskType); err != nil {
			r.logger.LogWarning(opSelectModuleTypeByTask, err, map[string]any{"userID": userTask.UserID, "labID": userTask.TaskID})
			return nil, err
		}
		return &taskType, nil
	}
	return nil, sql.ErrNoRows
}

func (r *studentRepository) SelectScore(ctx context.Context, userLab *model.UserLab) (*model.AssignedTaskByModule, error) {
	var score model.AssignedTaskByModule
	rows, err := r.conn.NamedQueryContext(ctx, selectScore, userLab)
	if err != nil {
		r.logger.LogWarning(opSelectScore, err, map[string]any{"userID": userLab.UserID, "labID": userLab.LabID})
		return nil, err
	}

	if rows.Next() {
		if err := rows.Scan(&score.Score); err != nil {
			r.logger.LogWarning(opSelectScore, err, map[string]any{"userID": userLab.UserID, "labID": userLab.LabID})
			return nil, err
		}
		return &score, nil
	}
	return nil, sql.ErrNoRows
}

func (r *studentRepository) BeginLab(ctx context.Context, userLab *model.UserLab) (*model.UserLab, error) {
	rows, err := r.conn.NamedQueryContext(ctx, beginLab, userLab)
	if err != nil {
		r.logger.LogWarning(opBeginLab, err, map[string]any{"userID": userLab.UserID, "labID": userLab.LabID})
		return nil, err
	}

	if rows.Next() {
		if err := rows.Scan(&userLab.LabID); err != nil {
			r.logger.LogWarning(opBeginLab, err, map[string]any{"userID": userLab.UserID, "labID": userLab.LabID})
			return nil, err
		}
		return userLab, nil
	}
	return nil, sql.ErrNoRows
}

func (r *studentRepository) FinishLab(ctx context.Context, userLab *model.UserLab) (*model.UserLab, error) {
	rows, err := r.conn.NamedQueryContext(ctx, finishLab, userLab)
	if err != nil {
		r.logger.LogWarning(opFinishLab, err, map[string]any{"userID": userLab.UserID, "labID": userLab.LabID})
		return nil, err
	}

	if rows.Next() {
		if err := rows.Scan(&userLab.LabID); err != nil {
			r.logger.LogWarning(opFinishLab, err, map[string]any{"userID": userLab.UserID, "labID": userLab.LabID})
			return nil, err
		}
		return userLab, nil
	}
	return nil, sql.ErrNoRows
}

func (r *studentRepository) SendAnswers(ctx context.Context, userLab *model.UserLabAnswer) (*model.UserLabAnswer, error) {
	rows, err := r.conn.NamedQueryContext(ctx, selectUserLabTask, userLab)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		if err := rows.StructScan(&userLab); err != nil {
			return nil, err
		}
	}

	var isActive bool
	err = r.conn.QueryRowxContext(ctx, checkLabActive, userLab.UserLabID).Scan(&isActive)
	if err != nil {
		r.logger.LogWarning(opFinishLab, err, map[string]any{"userID": userLab.UserID, "labID": userLab.LabID})
		return nil, err
	}
	if !isActive {
		return nil, ErrTimeExceeded
	}

	err = r.conn.QueryRowxContext(ctx, insertScore, userLab.UserLabID, userLab.TaskID, userLab.Answer, userLab.Score).Scan(&userLab.TaskID)
	if err != nil {
		return nil, err
	}
	if userLab.TaskID == -1 {
		return nil, ErrAnswerAlreadySent
	}

	return userLab, nil
}
