package common

import (
	"context"
	"github.com/labstack/echo/v4"
	"golang_graphs/internal/consts"
	"golang_graphs/internal/models"
	"log"
	"net/http"
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
		log.Println(consts.ErrorDescriptions[http.StatusBadRequest], err)
		return ctx.JSON(http.StatusBadRequest, models.BadRequestResponse{ErrorMsg: err.Error()})
	}

	ctxBack := context.Background()

	response, err := h.ctrl.SendAnswers(ctxBack, request)
	if err != nil {
		log.Println(consts.ErrorDescriptions[http.StatusBadRequest], err)
		return ctx.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse{ErrorMsg: err.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}
