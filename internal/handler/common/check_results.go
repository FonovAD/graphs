package common

import (
	"context"
	"github.com/labstack/echo/v4"
	"golang_graphs/internal/model"
	"net/http"
)

// CheckResults godoc
// @Summary      CheckResults
// @Description  CheckResults
// @Produce      json
// @Success      200  {object}  rest_models.CheckResultsResponse
// @Failure      400  {object}  model.BadRequestResponse
// @Failure      500  {object}  model.InternalServerErrorResponse
// @Router       /check_results [get]
func (h *handler) CheckResults(ctx echo.Context) error {
	userID := ctx.Get("user_id").(int64)

	ctxBack := context.Background()

	response, err := h.ctrl.CheckResults(ctxBack, userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.InternalServerErrorResponse{ErrorMsg: err.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}
