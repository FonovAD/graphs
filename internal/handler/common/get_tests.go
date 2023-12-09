package common

import (
	"context"
	"github.com/labstack/echo/v4"
	"golang_graphs/internal/model"
	"net/http"
)

// GetTests godoc
// @Summary      GetTests
// @Description  GetTests
// @Produce      json
// @Success      200  {object}  rest_models.GetTestsResponse
// @Failure      400  {object}  model.BadRequestResponse
// @Failure      500  {object}  model.InternalServerErrorResponse
// @Router       /get_tests [get]
func (h *handler) GetTests(ctx echo.Context) error {
	ctxBack := context.Background()

	response, err := h.ctrl.GetTests(ctxBack)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.InternalServerErrorResponse{ErrorMsg: err.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}
