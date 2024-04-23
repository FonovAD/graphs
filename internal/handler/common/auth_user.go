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

// AuthUser godoc
// @Summary      AuthUser
// @Description  AuthUser
// @Accept       json
// @Produce      json
// @Param        AuthUser   body      models.AuthUserRequest  true "AuthUser"
// @Success      200  {object}  models.AuthUserResponse
// @Failure      400  {object}  models.BadRequestResponse
// @Failure      500  {object}  models.InternalServerErrorResponse
// @Router       /auth_user [post]
func (h *handler) AuthUser(ctx echo.Context) error {
	var request models.AuthUserRequest

	if err := ctx.Bind(&request); err != nil {
		log.Println(consts.ErrorDescriptions[http.StatusBadRequest], err)
		return ctx.JSON(http.StatusBadRequest, models.BadRequestResponse{ErrorMsg: err.Error()})
	}

	ctxBack := context.Background()

	response, err := h.ctrl.AuthUser(ctxBack, request)
	if err != nil {
		log.Println(consts.ErrorDescriptions[http.StatusBadRequest], err)
		return ctx.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse{ErrorMsg: err.Error()})
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = response.Token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	ctx.SetCookie(cookie)

	return ctx.JSON(http.StatusOK, response)
}
