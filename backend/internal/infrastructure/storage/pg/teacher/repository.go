package storage

import (
	"context"
	"database/sql"
	model "golang_graphs/backend/internal/domain/model"
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
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&user.ID); err != nil {
			r.logger.LogWarning(opCreateLab, err, user)
			return nil, err
		}
		return user, nil
	}

	return nil, sql.ErrNoRows

}

func (r *teacherRepository) GetModules(ctx context.Context) ([]model.Module, error) {
	var modules []model.Module
	err := r.conn.SelectContext(ctx, &modules, selectAllModules)
	if err != nil {
		r.logger.LogDebug(opGetModules, err, nil)
		return nil, err
	}

	return modules, nil
}

func (r *teacherRepository) CreateLab(ctx context.Context, lab *model.Lab) (*model.Lab, error) {
	rows, err := r.conn.NamedQueryContext(ctx, createLab, lab)
	if err != nil {
		r.logger.LogDebug(opCreateLab, err, lab)
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&lab.ID); err != nil {
			r.logger.LogWarning(opCreateLab, err, lab)
			return nil, err
		}
		return lab, nil
	}

	return nil, sql.ErrNoRows
}

func (r *teacherRepository) GetLabInfo(ctx context.Context, lab *model.Lab) (*model.Lab, error) {
	rows, err := r.conn.NamedQueryContext(ctx, selectLabInfo, lab)
	if err != nil {
		r.logger.LogDebug(opGetLabInfo, err, lab)
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.StructScan(lab); err != nil {
			r.logger.LogWarning(opGetLabInfo, err, lab)
			return nil, err
		}
		return lab, nil
	}

	return nil, sql.ErrNoRows
}

func (r *teacherRepository) RemoveUserLab(ctx context.Context, userLab *model.UserLab) (*model.UserLab, error) {
	rows, err := r.conn.NamedQueryContext(ctx, removeLabFromUserLab, userLab)
	if err != nil {
		r.logger.LogDebug(opRemoveUserLab, err, userLab)
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&userLab.LabID); err != nil {
			r.logger.LogWarning(opRemoveUserLab, err, userLab)
			return nil, err
		}
		return userLab, nil
	}

	return nil, sql.ErrNoRows
}

func (r *teacherRepository) UpdateLab(ctx context.Context, lab *model.Lab) error {
	_, err := r.conn.NamedExecContext(ctx, updateLabInfo, lab)
	if err != nil {
		r.logger.LogDebug(opUpdateLab, err, lab)
		return err
	}

	return nil
}

func (r *teacherRepository) InsertUserLab(ctx context.Context, userLab *model.UserLab) (*model.UserLab, error) {
	rows, err := r.conn.NamedQueryContext(ctx, insertUserLab, userLab)
	if err != nil {
		r.logger.LogDebug(opInsertUserLab, err, userLab)
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&userLab.UserLabID); err != nil {
			r.logger.LogWarning(opInsertUserLab, err, userLab)
			return nil, err
		}
		return userLab, nil
	}

	return nil, sql.ErrNoRows

}

func (r *teacherRepository) InsertModuleLab(ctx context.Context, moduleLab *model.ModuleLab) (*model.ModuleLab, error) {
	rows, err := r.conn.NamedQueryContext(ctx, addModuleToLab, moduleLab)
	if err != nil {
		r.logger.LogDebug(opInsertModuleLab, err, moduleLab)
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&moduleLab.ModuleLabID); err != nil {
			r.logger.LogWarning(opInsertModuleLab, err, moduleLab)
			return nil, err
		}
		return moduleLab, nil
	}

	return nil, ErrModuleInLabExists
}

func (r *teacherRepository) RemoveModuleFromLab(ctx context.Context, moduleLab *model.ModuleLab) (*model.ModuleLab, error) {
	rows, err := r.conn.NamedQueryContext(ctx, removeModuleFromLab, moduleLab)
	if err != nil {
		r.logger.LogDebug(opRemoveModuleFromLab, err, moduleLab)
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&moduleLab.ModuleLabID); err != nil {
			r.logger.LogWarning(opRemoveModuleFromLab, err, moduleLab)
			return nil, err
		}
		return moduleLab, nil
	}

	return nil, sql.ErrNoRows
}

func (r *teacherRepository) InsertLabToStudentGroup(ctx context.Context, userLab *model.UserLabGroup) (*model.UserLabGroup, error) {
	var studentsCount int
	err := r.conn.QueryRowContext(ctx, selectStudentsCountFromGroup, userLab.GroupID).Scan(&studentsCount)
	if err != nil {
		return nil, err
	}

	var availableTasksCount int
	err = r.conn.QueryRowContext(ctx, selectAvailableTasksCountByModule, userLab.LabID).Scan(&availableTasksCount)
	if err != nil {
		return nil, err
	}

	if availableTasksCount < studentsCount {
		return nil, ErrTasksLessThanStudents
	}

	rows, err := r.conn.NamedQueryContext(ctx, insertLabToStudentGroup, userLab)
	if err != nil {
		r.logger.LogDebug(opInsertLabStudentGroup, err, userLab)
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&userLab.LabID); err != nil {
			r.logger.LogWarning(opInsertLabStudentGroup, err, userLab)
			return nil, err
		}
		return userLab, nil
	}

	return nil, sql.ErrNoRows
}

func (r *teacherRepository) SelectNonExistingUserLabs(ctx context.Context, pagination model.Pagination) ([]model.Lab, error) {
	var labs []model.Lab

	err := r.conn.SelectContext(ctx, &labs, selectNonExistingUserLabs, pagination.Limit, pagination.Offset)
	if err != nil {
		r.logger.LogDebug(opSelectNonExistingUserLabs, err, nil)
		return nil, err
	}

	return labs, nil
}

func (r *teacherRepository) SelectExistingUserLabs(ctx context.Context) ([]model.UserLabWithInfo, error) {
	var userLabs []model.UserLabWithInfo

	err := r.conn.SelectContext(ctx, &userLabs, selectExistingUserLabs)
	if err != nil {
		r.logger.LogDebug(opSelectExistingUserLabs, err, nil)
		return nil, err
	}

	return userLabs, nil
}

func (r *teacherRepository) SelectModulesFromLab(ctx context.Context, lab *model.Lab) ([]model.ModulesInLab, error) {
	var modules []model.ModulesInLab
	err := r.conn.SelectContext(ctx, &modules, selectModulesFromLab, lab.ID)
	if err != nil {
		r.logger.LogDebug(opSelectModulesFromLab, err, lab.ID)
		return nil, err
	}

	return modules, nil
}

func (r *teacherRepository) SelectTeacher(ctx context.Context, user *model.User) (*model.Teacher, error) {
	var teacher model.Teacher
	rows, err := r.conn.NamedQueryContext(ctx, selectTeacher, user)
	if err != nil {
		r.logger.LogWarning(opSelectTeacher, err, user.ID)
		return nil, err
	}

	if rows.Next() {
		if err := rows.Scan(&teacher.ID); err != nil {
			r.logger.LogWarning(opSelectTeacher, err, user.ID)
			return nil, err
		}
		teacher.UserID = user.ID
		return &teacher, nil
	}
	r.logger.LogDebug(opSelectTeacher, nil, user.ID)
	return nil, sql.ErrNoRows
}

func (r *teacherRepository) SelectGroups(ctx context.Context) ([]model.Group, error) {
	var groups []model.Group
	err := r.conn.SelectContext(ctx, &groups, selectGroups)
	if err != nil {
		r.logger.LogWarning(opSelectGroups, err, nil)
		return nil, err
	}

	return groups, nil
}

func (r *teacherRepository) InsertTask(ctx context.Context, task *model.Task) (*model.Task, error) {
	rows, err := r.conn.NamedQueryContext(ctx, insertTask, task)
	if err != nil {
		r.logger.LogDebug(opCreateTask, err, task)
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&task.ID); err != nil {
			r.logger.LogWarning(opCreateTask, err, task)
			return nil, err
		}
		return task, nil
	}

	return nil, sql.ErrNoRows
}

func (r *teacherRepository) UpdateTask(ctx context.Context, task *model.Task) (*model.Task, error) {
	rows, err := r.conn.NamedQueryContext(ctx, updateTask, task)
	if err != nil {
		r.logger.LogDebug(opUpdateTask, err, task)
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&task.ID); err != nil {
			r.logger.LogWarning(opUpdateTask, err, task)
			return nil, err
		}
		return task, nil
	}

	return nil, sql.ErrNoRows
}

func (r *teacherRepository) GetTasksByModule(ctx context.Context, module *model.Module) ([]model.TaskByModule, error) {
	var tasks []model.TaskByModule
	err := r.conn.SelectContext(ctx, &tasks, selectTasksByModule, module.ModuleId)
	if err != nil {
		r.logger.LogDebug(opSelectTasksByModule, err, module.ModuleId)
		return nil, err
	}

	return tasks, nil
}
