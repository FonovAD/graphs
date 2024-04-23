package common

import (
	"context"
	"github.com/labstack/echo/v4"
	"golang_graphs/internal/consts"
	"golang_graphs/internal/models"
	"log"
	"net/http"
	"time"
)

// CreateUser godoc
// @Summary      CreateUser
// @Description  CreateUser
// @Accept       json
// @Produce      json
// @Param        CreateUser   body      models.CreateUserRequest  true "CreateUser"
// @Success      200  {object}  models.CreateUserResponse
// @Failure      400  {object}  models.BadRequestResponse
// @Failure      500  {object}  models.InternalServerErrorResponse
// @Router       /create_user [post]
func (h *handler) CreateUser(ctx echo.Context) error {
	var request models.CreateUserRequest

	if err := ctx.Bind(&request); err != nil {
		log.Println(consts.ErrorDescriptions[http.StatusBadRequest], err)
		return ctx.JSON(http.StatusBadRequest, models.BadRequestResponse{ErrorMsg: err.Error()})
	}

	ctxBack := context.Background()

	response, err := h.ctrl.CreateUser(ctxBack, request)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse{ErrorMsg: err.Error()})
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = response.Token
	cookie.Expires = time.Now().Add(48 * time.Hour)
	ctx.SetCookie(cookie)

	return ctx.JSON(http.StatusOK, response)
}
