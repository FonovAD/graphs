package common

import (
	"context"
	"github.com/labstack/echo/v4"
	"golang_graphs/internal/consts"
	"golang_graphs/internal/model"
	"golang_graphs/internal/rest_models"
	"log"
	"net/http"
	"time"
)

// AuthUser godoc
// @Summary      AuthUser
// @Description  AuthUser
// @Accept       json
// @Produce      json
// @Param        AuthUser   body      rest_models.AuthUserRequest  true "AuthUser"
// @Success      200  {object}  rest_models.AuthUserResponse
// @Failure      400  {object}  model.BadRequestResponse
// @Failure      500  {object}  model.InternalServerErrorResponse
// @Router       /auth_user [post]
func (h *handler) AuthUser(ctx echo.Context) error {
	var request rest_models.AuthUserRequest

	if err := ctx.Bind(&request); err != nil {
		log.Println(consts.ErrorDescriptions[http.StatusBadRequest], err)
		return ctx.JSON(http.StatusBadRequest, model.BadRequestResponse{ErrorMsg: err.Error()})
	}

	ctxBack := context.Background()

	response, err := h.ctrl.AuthUser(ctxBack, request.Email, request.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.InternalServerErrorResponse{ErrorMsg: err.Error()})
	}

	token, err := h.authService.CreateToken(response)
	if err != nil {
		log.Println(consts.ErrorDescriptions[http.StatusInternalServerError], err)
		return ctx.JSON(http.StatusInternalServerError, model.BadRequestResponse{ErrorMsg: err.Error()})
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	ctx.SetCookie(cookie)

	return ctx.JSON(http.StatusOK, rest_models.AuthUserResponse{Token: token})
}
