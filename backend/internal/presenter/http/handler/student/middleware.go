package handler

import (
	"errors"
	usecase "golang_graphs/backend/internal/usecase/teacher"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const (
	authorizationHeader = "Authorization"
)

func (h *studentHandler) StudentMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			const op = "internal.handler.AuthMiddleware"

			header := ctx.Request().Header.Get(authorizationHeader)
			if header == "" {
				return ctx.JSON(http.StatusUnauthorized, BadRequestResponse{ErrorMsg: ErrEmptyToken.Error()})
			}

			headerSplit := strings.Split(header, " ")
			if len(headerSplit) != 2 {
				return ctx.JSON(http.StatusUnauthorized, UnauthorizedResponse{ErrorMsg: ErrInvalidAuthHeader.Error()})
			}

			student, err := h.studentUseCase.AuthToken(ctx.Request().Context(), headerSplit[1])
			if err != nil {
				if errors.Is(err, usecase.ErrNoPermissions) {
					return ctx.JSON(http.StatusForbidden, ForbiddenResponse{ErrorMsg: err.Error()})
				}
				if errors.Is(err, usecase.ErrParseToken) {
					return ctx.JSON(http.StatusUnauthorized, UnauthorizedResponse{ErrorMsg: ErrInvalidToken.Error()})
				}
				logrus.WithFields(logrus.Fields{"event": op}).Error(err)

				return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
			}

			ctx.Set("studentID", student.StudentID)
			ctx.Set("userID", student.UserID)
			return next(ctx)
		}
	}
}
