package repository

import (
	"context"
	model "golang_graphs/backend/internal/domain/model"
)

type TeacherRepository interface {
	InsertUser(ctx context.Context, user *model.User) (*model.User, error)
	GetModules(ctx context.Context) ([]model.Module, error)
	CreateLab(ctx context.Context, lab *model.Lab) (*model.Lab, error)
	GetLabInfo(ctx context.Context, lab *model.Lab) (*model.Lab, error)
	RemoveUserLab(ctx context.Context, userLab *model.UserLab) (*model.UserLab, error)
	UpdateLab(ctx context.Context, lab *model.Lab) error
	InsertUserLab(ctx context.Context, userLab *model.UserLab) (*model.UserLab, error)
	InsertModuleLab(ctx context.Context, moduleLab *model.ModuleLab) (*model.ModuleLab, error)
	RemoveModuleFromLab(ctx context.Context, moduleLab *model.ModuleLab) (*model.ModuleLab, error)
	InsertLabToStudentGroup(ctx context.Context, userLab *model.UserLabGroup) (*model.UserLabGroup, error)
	SelectNonExistingUserLabs(ctx context.Context, pagination model.Pagination) ([]model.Lab, error)
	SelectExistingUserLabs(ctx context.Context) ([]model.UserLabWithInfo, error)
	SelectModulesFromLab(ctx context.Context, lab *model.Lab) ([]model.ModulesInLab, error)
	SelectTeacher(ctx context.Context, user *model.User) (*model.Teacher, error)
	SelectGroups(ctx context.Context) ([]model.Group, error)
	InsertTask(ctx context.Context, task *model.Task) (*model.Task, error)
	UpdateTask(ctx context.Context, task *model.Task) (*model.Task, error)
	GetTasksByModule(ctx context.Context, module *model.Module) ([]model.TaskByModule, error)
	GetGroupLabResults(ctx context.Context, group *model.Group) ([]model.GroupLabResult, error)
}
