package common

import (
	"context"
	"github.com/labstack/echo/v4"
	"golang_graphs/internal/models"
	"net/http"
)

// InsertTask godoc
// @Summary      InsertTask
// @Description  InsertTask
// @Accept       json
// @Produce      json
// @Param        InsertTest   body      models.InsertTaskRequest  true "InsertTest"
// @Success      200  {object}  models.InsertTaskResponse
// @Failure      400  {object}  models.BadRequestResponse
// @Failure      500  {object}  models.InternalServerErrorResponse
// @Router       /insert_task [post]
func (h *handler) InsertTask(ctx echo.Context) error {
	var request models.InsertTaskRequest

	if err := ctx.Bind(&request); err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, models.BadRequestResponse{ErrorMsg: err.Error()})
	}

	ctxBack := context.Background()
	response, err := h.ctrl.InsertTask(ctxBack, request)
	if err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse{ErrorMsg: err.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}
