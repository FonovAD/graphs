package common

import (
	"context"
	"github.com/labstack/echo/v4"
	"golang_graphs/backend/internal/models"
	"net/http"
)

// GetTests godoc
// @Summary      GetTests
// @Description  GetTests
// @Produce      json
// @Success      200  {object}  models.GetTestsResponse
// @Failure      400  {object}  models.BadRequestResponse
// @Failure      500  {object}  models.InternalServerErrorResponse
// @Router       /get_tests [get]
func (h *handler) GetTests(ctx echo.Context) error {
	ctxBack := context.Background()

	response, err := h.ctrl.GetTests(ctxBack)
	if err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse{ErrorMsg: err.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}
