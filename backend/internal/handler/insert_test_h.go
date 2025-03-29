package handler

import (
	"context"
	"golang_graphs/backend/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// InsertTest godoc
// @Summary      InsertTest
// @Description  InsertTest
// @Accept       json
// @Produce      json
// @Param        InsertTest   body      models.InsertTestRequest  true "InsertTest"
// @Success      200  {object}  models.InsertTestResponse
// @Failure      400  {object}  models.BadRequestResponse
// @Failure      500  {object}  models.InternalServerErrorResponse
// @Router       /insert_test [post]
func (h *handler) InsertTest(ctx echo.Context) error {
	var request models.InsertTestRequest

	if err := ctx.Bind(&request); err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, models.BadRequestResponse{ErrorMsg: err.Error()})
	}

	ctxBack := context.Background()
	response, err := h.ctrl.InsertTest(ctxBack, request)
	if err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}
