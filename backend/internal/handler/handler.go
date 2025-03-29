package handler

import (
	"golang_graphs/backend/internal/controller"

	"github.com/labstack/echo/v4"
)

// New constructor for Handler, user for code generation in wire
func New(ctrl controller.Controller) Handler {
	return &handler{ctrl: ctrl}
}

type Handler interface {
	GetTests(ctx echo.Context) error
	GetTasksFromTest(ctx echo.Context) error
	CreateUser(ctx echo.Context) error
	CheckResults(ctx echo.Context) error
	AuthUser(ctx echo.Context) error
	SendAnswers(ctx echo.Context) error
	InsertTask(ctx echo.Context) error
	InsertTest(ctx echo.Context) error
	Ping(ctx echo.Context) error
	AuthMiddleware() echo.MiddlewareFunc
}

// Handler struct for declaring api methods
type handler struct {
	ctrl controller.Controller
}
