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
// @Param        CheckResultsRequest   body      models.CheckResultsRequest  true "CheckResultsRequest"
// @Success      200  {object}  models.CheckResultsResponse
// @Failure      400  {object}  models.BadRequestResponse
// @Failure      500  {object}  models.InternalServerErrorResponse
// @Router       /check_results [post]
func (h *handler) CheckResults(ctx echo.Context) error {
	var request models.CheckResultsRequest
	if err := ctx.Bind(&request); err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, models.BadRequestResponse{ErrorMsg: err.Error()})
	}

	ctxBack := context.Background()
	response, err := h.ctrl.CheckResults(ctxBack, request)
	if err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse{ErrorMsg: err.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}
