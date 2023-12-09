package task

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"golang_graphs/internal/consts"
	"golang_graphs/internal/controller/controller_task"
	"golang_graphs/internal/model"
	"log"
	"net/http"
)

// Handler struct for declaring api methods
type handler struct {
	ctrl controller_task.Controller
}

// NewHandler constructor for Handler, user for code generation in wire
func NewHandler(ctrl controller_task.Controller) Handler {
	return &handler{ctrl: ctrl}
}

type Handler interface {
	TaskComponents(ctx echo.Context) error
	TaskIsEulerUndirected(ctx echo.Context) error
	TaskIsBipartition(ctx echo.Context) error
}

func (h *handler) TaskComponents(ctx echo.Context) error {
	var request model.Graph

	if err := ctx.Bind(&request); err != nil {
		log.Println(consts.ErrorDescriptions[http.StatusBadRequest], err)
		return ctx.JSON(http.StatusBadRequest, model.BadRequestResponse{})
	}

	// fmt.Println("graph", request)
	// fmt.Println("graph links", request.Links)
	// fmt.Println("graph nodes", request.Nodes)

	ctxBack := context.Background()

	components, err := h.ctrl.TaskComponents(ctxBack, request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.BadRequestResponse{})
	}

	fmt.Println("components", components)

	return ctx.JSON(http.StatusOK, model.Component{Component: components})
}

func (h *handler) TaskIsEulerUndirected(ctx echo.Context) error {
	var request model.Graph

	if err := ctx.Bind(&request); err != nil {
		log.Println(consts.ErrorDescriptions[http.StatusBadRequest], err)
		return ctx.JSON(http.StatusBadRequest, model.BadRequestResponse{})
	}

	// fmt.Println("graph", request)
	// fmt.Println("graph links", request.Links)
	// fmt.Println("graph nodes", request.Nodes)

	ctxBack := context.Background()

	isEuler, err := h.ctrl.TaskIsEulerUndirected(ctxBack, request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.BadRequestResponse{})
	}

	fmt.Println("isEuler", isEuler)

	return ctx.JSON(http.StatusOK, model.IsEuler{IsEuler: isEuler})
}

func (h *handler) TaskIsBipartition(ctx echo.Context) error {
	var request model.Graph

	if err := ctx.Bind(&request); err != nil {
		log.Println(consts.ErrorDescriptions[http.StatusBadRequest], err)
		return ctx.JSON(http.StatusBadRequest, model.BadRequestResponse{})
	}

	// fmt.Println("graph", request)
	// fmt.Println("graph links", request.Links)
	// fmt.Println("graph nodes", request.Nodes)

	ctxBack := context.Background()

	isBipartition, err := h.ctrl.TaskIsBipartition(ctxBack, request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.BadRequestResponse{})
	}

	fmt.Println("isBipartition", isBipartition)

	return ctx.JSON(http.StatusOK, model.IsBipartition{IsBipartition: isBipartition})
}
