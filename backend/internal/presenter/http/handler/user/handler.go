package userhandler

import (
	"context"
	usecase "golang_graphs/backend/internal/usecase/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	AuthUser(ctx echo.Context) error
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(u usecase.UserUseCase) UserHandler {
	return &userHandler{u}
}

func (h *userHandler) AuthUser(ctx echo.Context) error {
	var request AuthUserRequest
	if err := ctx.Bind(&request); err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}

	ctxBack := context.Background()

	authUserDTO := usecase.AuthUserDTO{Email: request.Email, Password: request.Password}
	response, err := h.userUseCase.AuthUser(ctxBack, authUserDTO)
	if err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: err.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}
