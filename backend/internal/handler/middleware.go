package handler

import (
	"golang_graphs/backend/internal/models"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const (
	authorizationHeader = "Authorization"
)

func (h *handler) AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			const op = "internal.handler.AuthMiddleware"

			header := ctx.Request().Header.Get(authorizationHeader)
			if header == "" {
				return ctx.JSON(http.StatusUnauthorized, models.BadRequestResponse{ErrorMsg: ErrEmptyToken.Error()})
			}

			headerSplit := strings.Split(header, " ")
			if len(headerSplit) != 2 {
				return ctx.JSON(http.StatusUnauthorized, models.InternalServerErrorResponse{ErrorMsg: ErrInvalidAuthHeader.Error()})
			}

			user, err := h.ctrl.AuthToken(ctx.Request().Context(), headerSplit[1])
			if err != nil {
				logrus.WithFields(logrus.Fields{"event": op}).Error(err)

				return ctx.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
			}

			ctx.Set("user", user)
			return next(ctx)
		}
	}
}
