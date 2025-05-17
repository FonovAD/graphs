package interactor

import (
	"golang_graphs/backend/internal/logger"
	"golang_graphs/backend/internal/presenter/http/handler"
	studenthandler "golang_graphs/backend/internal/presenter/http/handler/student"
	teacherhandler "golang_graphs/backend/internal/presenter/http/handler/teacher"
	userhandler "golang_graphs/backend/internal/presenter/http/handler/user"

	"github.com/jmoiron/sqlx"
)

type Interactor interface {
	NewAppHandler() handler.AppHandler
}

type interactor struct {
	conn   *sqlx.DB
	logger logger.Logger
}

func NewInteractor(conn *sqlx.DB, logger logger.Logger) Interactor {
	return &interactor{conn: conn, logger: logger}
}

type appHandler struct {
	userhandler.UserHandler
	teacherhandler.TeacherHandler
	studenthandler.StudentHandler
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.UserHandler = i.NewUserHandler()
	appHandler.TeacherHandler = i.NewTeacherHandler()
	appHandler.StudentHandler = i.NewStudentHandler()
	return appHandler
}
