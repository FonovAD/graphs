package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) Ping(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "pong")
}
