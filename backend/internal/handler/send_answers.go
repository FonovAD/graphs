package handler

import (
	"context"
	"golang_graphs/backend/internal/dto"
	"golang_graphs/backend/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// SendAnswers godoc
// @Summary      SendAnswers
// @Description  SendAnswers
// @Accept       json
// @Produce      json
// @Param        SendAnswers   body      models.SendAnswersRequest  true "SendAnswers"
// @Success      200  {object}  models.SendAnswersResponse
// @Failure      400  {object}  models.BadRequestResponse
// @Failure      500  {object}  models.InternalServerErrorResponse
// @Router       /send_answers [post]
func (h *handler) SendAnswers(ctx echo.Context) error {
	var request models.SendAnswersRequest

	if err := ctx.Bind(&request); err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, models.BadRequestResponse{ErrorMsg: err.Error()})
	}

	user, ok := ctx.Get("user").(dto.User)

	if !ok {
		return ctx.JSON(http.StatusBadRequest, models.InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	ctxBack := context.Background()
	response, err := h.ctrl.SendAnswers(ctxBack, user, request)

	if err != nil && response.TaskType == -1 {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}
