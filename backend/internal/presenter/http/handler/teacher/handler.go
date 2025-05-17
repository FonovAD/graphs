package handler

import (
	"context"
	"errors"
	usecase "golang_graphs/backend/internal/usecase/teacher"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TeacherHandler interface {
	CreateUser(ctx echo.Context) error
}

type teacherHandler struct {
	teacherUseCase usecase.TeacherUseCase
}

func NewTeacherHandler(u usecase.TeacherUseCase) TeacherHandler {
	return &teacherHandler{u}
}

func (h *teacherHandler) CreateUser(ctx echo.Context) error {
	var request CreateUserRequest
	if err := ctx.Bind(&request); err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}

	ctxBack := context.Background()

	createUserDTO := usecase.CreateUserDTO{
		Email:      request.Email,
		Password:   request.Password,
		FirstName:  request.FirstName,
		LastName:   request.LastName,
		FatherName: request.FatherName,
	}
	response, err := h.teacherUseCase.CreateUser(ctxBack, createUserDTO)
	if err != nil && (errors.Is(err, usecase.ErrShortPassword) ||
		errors.Is(err, usecase.ErrShortFirstname) ||
		errors.Is(err, usecase.ErrShortLastname)) {
		ctx.Set("error", err.Error())

		return ctx.JSON(http.StatusInternalServerError, BadRequestResponse{ErrorMsg: err.Error()})
	}

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}
