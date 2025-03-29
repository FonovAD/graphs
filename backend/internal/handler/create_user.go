package handler

import (
	"context"
	"errors"
	"golang_graphs/backend/internal/controller"
	"golang_graphs/backend/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateUser godoc
// @Summary      CreateUser
// @Description  CreateUser
// @Accept       json
// @Produce      json
// @Param        CreateUser   body      models.CreateUserRequest  true "CreateUser"
// @Success      200  {object}  models.CreateUserResponse
// @Failure      400  {object}  models.BadRequestResponse
// @Failure      500  {object}  models.InternalServerErrorResponse
// @Router       /create_user [post]
func (h *handler) CreateUser(ctx echo.Context) error {
	var request models.CreateUserRequest
	if err := ctx.Bind(&request); err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, models.BadRequestResponse{ErrorMsg: err.Error()})
	}

	ctxBack := context.Background()
	response, err := h.ctrl.CreateUser(ctxBack, request)
	if err != nil && (errors.Is(err, controller.ErrShortPassword) ||
		errors.Is(err, controller.ErrShortFirstname) ||
		errors.Is(err, controller.ErrShortLastname)) {
		ctx.Set("error", err.Error())

		return ctx.JSON(http.StatusInternalServerError, models.BadRequestResponse{ErrorMsg: err.Error()})
	}

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}
