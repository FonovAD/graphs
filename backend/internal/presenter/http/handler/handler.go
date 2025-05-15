package handler

import (
	studenthandler "golang_graphs/backend/internal/presenter/http/handler/student"
	teacherhandler "golang_graphs/backend/internal/presenter/http/handler/teacher"
	userhandler "golang_graphs/backend/internal/presenter/http/handler/user"
)

type AppHandler interface {
	userhandler.UserHandler
	teacherhandler.TeacherHandler
	studenthandler.StudentHandler
}
