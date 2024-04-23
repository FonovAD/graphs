package common

import (
	"context"
	"github.com/labstack/echo/v4"
	"golang_graphs/internal/models"
	"net/http"
)

// CheckResults godoc
// @Summary      CheckResults
// @Description  CheckResults
// @Produce      json
// @Success      200  {object}  models.CheckResultsResponse
// @Failure      400  {object}  models.BadRequestResponse
// @Failure      500  {object}  models.InternalServerErrorResponse
// @Router       /check_results [get]
func (h *handler) CheckResults(ctx echo.Context) error {
	userID, ok := ctx.Get("user_id").(int64)
	if !ok {
		return ctx.JSON(http.StatusBadRequest, models.BadRequestResponse{ErrorMsg: "invalid user_id"})
	}

	ctxBack := context.Background()

	response, err := h.ctrl.CheckResults(ctxBack, userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse{ErrorMsg: err.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}
