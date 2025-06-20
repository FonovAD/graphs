package handler

import (
	usecase "golang_graphs/backend/internal/usecase/student"
	// "context"
	// "golang_graphs/backend/internal/dto"
	// "golang_graphs/backend/internal/models"
	// "net/http"

	// "github.com/labstack/echo/v4"
)

type StudentHandler interface{}

type studentHandler struct {
	studentUseCase usecase.StudentUseCase
}

func NewStudentHandler(u usecase.StudentUseCase) StudentHandler {
	return &studentHandler{u}
}

// func (h *handler) CheckResults(ctx echo.Context) error {
// 	var request models.CheckResultsRequest
// 	if err := ctx.Bind(&request); err != nil {
// 		ctx.Set("error", err.Error())
// 		return ctx.JSON(http.StatusBadRequest, models.BadRequestResponse{ErrorMsg: err.Error()})
// 	}

// 	user, ok := ctx.Get("user").(dto.User)

// 	if !ok {
// 		return ctx.JSON(http.StatusBadRequest, models.InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
// 	}

// 	ctxBack := context.Background()
// 	response, err := h.ctrl.CheckResults(ctxBack, user, request)
// 	if err != nil {
// 		ctx.Set("error", err.Error())
// 		return ctx.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
// 	}

// 	return ctx.JSON(http.StatusOK, response)
// }



// func (h *handler) SendAnswers(ctx echo.Context) error {
// 	var request models.SendAnswersRequest

// 	if err := ctx.Bind(&request); err != nil {
// 		ctx.Set("error", err.Error())
// 		return ctx.JSON(http.StatusBadRequest, models.BadRequestResponse{ErrorMsg: err.Error()})
// 	}

// 	user, ok := ctx.Get("user").(dto.User)

// 	if !ok {
// 		return ctx.JSON(http.StatusBadRequest, models.InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
// 	}

// 	ctxBack := context.Background()
// 	response, err := h.ctrl.SendAnswers(ctxBack, user, request)

// 	if err != nil && response.TaskType == -1 {
// 		ctx.Set("error", err.Error())
// 		return ctx.JSON(http.StatusBadRequest, response)
// 	}
// 	if err != nil {
// 		ctx.Set("error", err.Error())
// 		return ctx.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
// 	}

// 	return ctx.JSON(http.StatusOK, response)
// }
