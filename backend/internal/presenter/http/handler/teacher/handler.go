package handler

import (
	"context"
	"errors"
	storage "golang_graphs/backend/internal/infrastructure/storage/pg/teacher"
	usecase "golang_graphs/backend/internal/usecase/teacher"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TeacherHandler interface {
	CreateUser(ctx echo.Context) error
	GetModules(ctx echo.Context) error
	CreateLab(ctx echo.Context) error
	GetLabInfo(ctx echo.Context) error
	RemoveUserLab(ctx echo.Context) error
	UpdateLabInfo(ctx echo.Context) error
	AssignLab(ctx echo.Context) error
	AssignLabGroup(ctx echo.Context) error
	AddModuleLab(ctx echo.Context) error
	RemoveModuleLab(ctx echo.Context) error
	GetNonAssignedLabs(ctx echo.Context) error
	GetAssignedLabs(ctx echo.Context) error
	GetLabModules(ctx echo.Context) error
	TeacherMiddleware() echo.MiddlewareFunc
	GetGroups(ctx echo.Context) error
	CreateTask(ctx echo.Context) error
	UpdateTask(ctx echo.Context) error
	GetTasksByModule(ctx echo.Context) error
	GetLabResults(ctx echo.Context) error
}

type teacherHandler struct {
	teacherUseCase usecase.TeacherUseCase
}

func NewTeacherHandler(u usecase.TeacherUseCase) TeacherHandler {
	return &teacherHandler{u}
}

func (h *teacherHandler) CreateUser(ctx echo.Context) error {
	var request usecase.CreateUserDTOIn
	if err := ctx.Bind(&request); err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}

	ctxBack := context.Background()

	response, err := h.teacherUseCase.CreateStudent(ctxBack, &request)
	if err != nil {
		if errors.Is(err, usecase.ErrShortPassword) || errors.Is(err, usecase.ErrShortFirstname) || errors.Is(err, usecase.ErrShortLastname) {

			ctx.Set("error", err.Error())

			return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
		}

		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *teacherHandler) GetModules(ctx echo.Context) error {
	ctxBack := context.Background()

	response, err := h.teacherUseCase.GetModules(ctxBack)
	if err != nil {
		ctx.Set("error", err.Error())

		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *teacherHandler) CreateLab(ctx echo.Context) error {
	var request usecase.CreateLabDTOIn
	if err := ctx.Bind(&request); err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}
	request.TeacherID = ctx.Get("teacherID").(int64)

	ctxBack := context.Background()
	response, err := h.teacherUseCase.CreateLab(ctxBack, &request)
	if err != nil {
		ctx.Set("error", err.Error())

		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *teacherHandler) GetLabInfo(ctx echo.Context) error {
	var request usecase.GetLabInfoDTOIn
	if err := ctx.Bind(&request); err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}

	ctxBack := context.Background()
	response, err := h.teacherUseCase.GetLabInfo(ctxBack, &request)
	if err != nil {
		ctx.Set("error", err.Error())

		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *teacherHandler) RemoveUserLab(ctx echo.Context) error {
	var request usecase.RemoveUserLabDTOIn
	if err := ctx.Bind(&request); err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}

	ctxBack := context.Background()
	response, err := h.teacherUseCase.RemoveUserLab(ctxBack, &request)
	if err != nil {
		ctx.Set("error", err.Error())

		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *teacherHandler) UpdateLabInfo(ctx echo.Context) error {
	var request usecase.UpdateLabInfoDTOIn
	if err := ctx.Bind(&request); err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}

	ctxBack := context.Background()
	err := h.teacherUseCase.UpdateLabInfo(ctxBack, &request)
	if err != nil {
		ctx.Set("error", err.Error())

		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusOK, nil)
}

func (h *teacherHandler) AssignLab(ctx echo.Context) error {
	var request usecase.AssignLabDTOIn
	if err := ctx.Bind(&request); err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}
	request.AssigneID = ctx.Get("teacherID").(int64)

	ctxBack := context.Background()
	response, err := h.teacherUseCase.AssignLab(ctxBack, &request)
	if err != nil {
		ctx.Set("error", err.Error())

		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *teacherHandler) AssignLabGroup(ctx echo.Context) error {
	var request usecase.AssignLabGroupDTOIn
	if err := ctx.Bind(&request); err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}
	request.AssigneID = ctx.Get("teacherID").(int64)

	ctxBack := context.Background()
	response, err := h.teacherUseCase.AssignLabGroup(ctxBack, &request)
	if err != nil {
		ctx.Set("error", err.Error())

		if errors.Is(err, storage.ErrTasksLessThanStudents) {
			return ctx.JSON(http.StatusInternalServerError, BadRequestResponse{ErrorMsg: err.Error()})
		}

		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *teacherHandler) AddModuleLab(ctx echo.Context) error {
	var request usecase.AddModuleLabDTOIn
	if err := ctx.Bind(&request); err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}

	ctxBack := context.Background()
	response, err := h.teacherUseCase.AddModuleLab(ctxBack, &request)
	if err != nil {
		ctx.Set("error", err.Error())

		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *teacherHandler) RemoveModuleLab(ctx echo.Context) error {
	var request usecase.GetLabInfoDTOIn
	if err := ctx.Bind(&request); err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}

	ctxBack := context.Background()
	response, err := h.teacherUseCase.GetLabInfo(ctxBack, &request)
	if err != nil {
		ctx.Set("error", err.Error())

		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *teacherHandler) GetNonAssignedLabs(ctx echo.Context) error {
	var request usecase.GetNonAssignedLabsDTOIn
	if err := ctx.Bind(&request); err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}

	ctxBack := context.Background()
	response, err := h.teacherUseCase.GetNonAssignedLabs(ctxBack, &request)
	if err != nil {
		ctx.Set("error", err.Error())

		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *teacherHandler) GetAssignedLabs(ctx echo.Context) error {
	ctxBack := context.Background()
	response, err := h.teacherUseCase.GetAssignedLabs(ctxBack)
	if err != nil {
		ctx.Set("error", err.Error())

		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *teacherHandler) GetLabModules(ctx echo.Context) error {
	var request usecase.GetLabModulesDTOIn
	if err := ctx.Bind(&request); err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}

	ctxBack := context.Background()
	response, err := h.teacherUseCase.GetLabModules(ctxBack, &request)
	if err != nil {
		ctx.Set("error", err.Error())

		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *teacherHandler) GetGroups(ctx echo.Context) error {
	ctxBack := context.Background()

	response, err := h.teacherUseCase.GetGroups(ctxBack)
	if err != nil {
		ctx.Set("error", err.Error())

		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *teacherHandler) CreateTask(ctx echo.Context) error {
	var request usecase.CreateTaskDTOIn
	if err := ctx.Bind(&request); err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}

	ctxBack := context.Background()
	response, err := h.teacherUseCase.CreateTask(ctxBack, &request)
	if err != nil {
		ctx.Set("error", err.Error())

		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *teacherHandler) UpdateTask(ctx echo.Context) error {
	var request usecase.CreateTaskDTOIn
	if err := ctx.Bind(&request); err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}

	ctxBack := context.Background()
	response, err := h.teacherUseCase.UpdateTask(ctxBack, &request)
	if err != nil {
		ctx.Set("error", err.Error())

		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *teacherHandler) GetTasksByModule(ctx echo.Context) error {
	var request usecase.GetTasksByModuleDTOIn
	if err := ctx.Bind(&request); err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}

	ctxBack := context.Background()
	response, err := h.teacherUseCase.GetTasksByModule(ctxBack, &request)
	if err != nil {
		ctx.Set("error", err.Error())

		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *teacherHandler) GetLabResults(ctx echo.Context) error {
	var request usecase.GetLabResultsDTOIn
	if err := ctx.Bind(&request); err != nil {
		ctx.Set("error", err.Error())
		return ctx.JSON(http.StatusBadRequest, BadRequestResponse{ErrorMsg: err.Error()})
	}

	ctxBack := context.Background()
	response, err := h.teacherUseCase.GetLabResults(ctxBack, &request)
	if err != nil {
		ctx.Set("error", err.Error())

		return ctx.JSON(http.StatusInternalServerError, InternalServerErrorResponse{ErrorMsg: ErrInternalServer.Error()})
	}

	return ctx.JSON(http.StatusOK, response)
}
