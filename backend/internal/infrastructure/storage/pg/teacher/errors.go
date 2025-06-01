package storage

import "errors"

var (
	ErrModuleInLabExists     = errors.New("module in lab already exists")
	ErrTasksLessThanStudents = errors.New("tasks count less than students count")
)
