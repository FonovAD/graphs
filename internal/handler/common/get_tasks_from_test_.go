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

// GetTasksFromTest godoc
// @Summary      GetTasksFromTest
// @Description  GetTasksFromTest
// @Accept       json
// @Produce      json
// @Param        GetTasksFromTest   body      rest_models.GetTasksFromTestsRequest  true "GetTasksFromTest"
// @Success      200  {object}  rest_models.GetTasksFromTestsResponse
// @Failure      400  {object}  model.BadRequestResponse
// @Failure      500  {object}  model.InternalServerErrorResponse
// @Router       /get_tasks_from_test [post]
func (h *handler) GetTasksFromTest(ctx echo.Context) error {
	var request rest_models.GetTasksFromTestsRequest

	if err := ctx.Bind(&request); err != nil {
		log.Println(consts.ErrorDescriptions[http.StatusBadRequest], err)
		return ctx.JSON(http.StatusBadRequest, model.BadRequestResponse{ErrorMsg: err.Error()})
	}

	ctxBack := context.Background()

	tasks, err := h.ctrl.GetTasksFromTest(ctxBack, request.TestID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.InternalServerErrorResponse{ErrorMsg: err.Error()})
	}

	// удаляю ответы
	// TODO Очень грубый костыль, необходимо перенести эту логику в котроллер (создать новый метод), так как GetTasksFromTest уже используется в другом месте
	for i := 0; i < len(tasks); i++ {
		tasks[i].Answer = ""
	}

	return ctx.JSON(http.StatusOK, rest_models.GetTasksFromTestsResponse{Tasks: tasks})
}
