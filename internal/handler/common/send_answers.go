package common

import (
	"context"
	"github.com/labstack/echo/v4"
	"golang_graphs/internal/consts"
	"golang_graphs/internal/model"
	"golang_graphs/internal/rest_models"
	"log"
	"net/http"
)

// SendAnswers godoc
// @Summary      SendAnswers
// @Description  SendAnswers
// @Accept       json
// @Produce      json
// @Param        SendAnswers   body      rest_models.SendAnswersRequest  true "SendAnswers"
// @Success      200  {object}  rest_models.SendAnswersResponse
// @Failure      400  {object}  model.BadRequestResponse
// @Failure      500  {object}  model.InternalServerErrorResponse
// @Router       /send_answers [post]
func (h *handler) SendAnswers(ctx echo.Context) error {
	var request rest_models.SendAnswersRequest

	if err := ctx.Bind(&request); err != nil {
		log.Println(consts.ErrorDescriptions[http.StatusBadRequest], err)
		return ctx.JSON(http.StatusBadRequest, model.BadRequestResponse{ErrorMsg: err.Error()})
	}

	ctxBack := context.Background()

	result, err := h.ctrl.SendAnswers(ctxBack, request.Answers, request.TestID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.InternalServerErrorResponse{ErrorMsg: err.Error()})
	}

	return ctx.JSON(http.StatusOK, rest_models.SendAnswersResponse{Result: result})
}
