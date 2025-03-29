import sys
import re


def camel_case_to_snake_case(name):
    name = re.sub(r'(?<!^)(?=[A-Z])', '_', name).lower()
    return name


def create_file_name(name, path):
    return path + camel_case_to_snake_case(name) + ".go"


def replace(name):
    function_name = "{TEMPLATE}"
    path_name = "{TEMPLATE_PATH}"

    template = """
package common

import (
    "context"
    "github.com/labstack/echo/v4"
    "golang_graphs/internal/consts"
    "golang_graphs/internal/model"
    "log"
    "net/http"
)

type {TEMPLATE}Request struct {
}

type {TEMPLATE}Response struct {
}

// {TEMPLATE} godoc
// @Summary      {TEMPLATE}
// @Description  {TEMPLATE}
// @Accept       json
// @Produce      json
// @Param        {TEMPLATE}   body      {TEMPLATE}Request  true "{TEMPLATE}"
// @Success      200  {object}  {TEMPLATE}Response
// @Failure      400  {object}  model.BadRequestResponse
// @Failure      500  {object}  model.InternalServerErrorResponse
// @Router       /{TEMPLATE_PATH} [post]
func (h *handler) {TEMPLATE}(ctx echo.Context) error {
    var request {TEMPLATE}Request

    if err := ctx.Bind(&request); err != nil {
        log.Println(consts.ErrorDescriptions[http.StatusBadRequest], err)
        return ctx.JSON(http.StatusBadRequest, model.BadRequestResponse{ErrorMsg: err.Error()})
    }

    ctxBack := context.Background()

    response, err := h.ctrl.{TEMPLATE}(ctxBack, request)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, model.InternalServerErrorResponse{ErrorMsg: err.Error()})
    }

    return ctx.JSON(http.StatusOK, response)
}
    """

    template = template.strip()

    file_data = str.replace(template, function_name, name)
    file_data = str.replace(file_data, path_name, camel_case_to_snake_case(name))

    return file_data


### args[1] - название функции
### args[2] - путь до файла
def main():
    args = sys.argv

    if len(args) <= 2:
        print("error: need args")
        return

    file_data = replace(args[1])
    file_name = create_file_name(args[1], args[2])

    file = open(file_name, 'w')
    file.write(file_data)


if __name__ == "__main__":
    main()
