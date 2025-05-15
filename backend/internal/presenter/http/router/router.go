package router

import (
	"golang_graphs/backend/internal/presenter/http/handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.GET("/ping", Ping)

	userRouter := e.Group("/api/v1")
	userRouter.POST("/auth_user", h.AuthUser)

	teacherRouter := e.Group("/api/v1")
	teacherRouter.POST("/create_user", h.CreateUser)
}

func Ping(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "pong")
}
