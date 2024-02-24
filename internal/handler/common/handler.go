package common

import (
	"github.com/labstack/echo/v4"
	"golang_graphs/internal/controller"
	"golang_graphs/pkg/auth"
)

// New constructor for Handler, user for code generation in wire
func New(ctrl controller.Controller, authService auth.Service) Handler {
	return &handler{ctrl: ctrl, authService: authService}
}

type Handler interface {
	GetTests(ctx echo.Context) error
	GetTasksFromTest(ctx echo.Context) error
	CreateUser(ctx echo.Context) error
	CheckResults(ctx echo.Context) error
	AuthUser(ctx echo.Context) error
	SendAnswers(ctx echo.Context) error
}

// Handler struct for declaring api methods
type handler struct {
	ctrl        controller.Controller
	authService auth.Service
}
