package handler

import (
	"context"
	"golang_graphs/backend/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetTasksFromTest godoc
// @Summary      GetTasksFromTest
// @Description  GetTasksFromTest
// @Accept       json
// @Produce      json
// @Param        GetTasksFromTest   body      models.GetTasksFromTestsRequest  true "GetTasksFromTest"
// @Success      200  {object}  models.GetTasksFromTestsResponse
// @Failure      400  {object}  models.BadRequestResponse
// @Failure      500  {object}  models.InternalServerErrorResponse
// @Router       /get_tasks_from_test [post]
func (h *handler) GetTasksFromTest(ctx echo.Context) error {
	var request models.GetTasksFromTestsRequest
	if err := ctx.Bind(&request); err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, models.BadRequestResponse{ErrorMsg: err.Error()})
	}

	ctxBack := context.Background()
	response, err := h.ctrl.GetTasksFromTest(ctxBack, request)
	if err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}
