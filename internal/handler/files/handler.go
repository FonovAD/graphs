package files

import (
	"github.com/labstack/echo/v4"
)

// handler struct for declaring api methods
type handler struct {
	rootPath string
}

type Handler interface {
	GetCSS(ctx echo.Context) error
	GetJSD3(ctx echo.Context) error
	GetJS(ctx echo.Context) error
	Favicon(ctx echo.Context) error
	GetScript(ctx echo.Context) error
	GetCSSLogin(ctx echo.Context) error
}

// NewHandler constructor for handler, user for code generation in wire
func NewHandler(rootPath string) Handler {
	return &handler{rootPath: rootPath}
}

func (h *handler) GetCSS(ctx echo.Context) error {
	return ctx.File(h.rootPath + "/internal/view/graphPlayground/css/style.css")
}

func (h *handler) GetJSD3(ctx echo.Context) error {
	return ctx.File(h.rootPath + "/internal/view/graphPlayground/js/d3.v5.min.js")
}

func (h *handler) GetJS(ctx echo.Context) error {
	return ctx.File(h.rootPath + "/internal/view/graphPlayground/js/script.js")
}

func (h *handler) Favicon(ctx echo.Context) error {
	return ctx.File(h.rootPath + "/internal/view/graphPlayground/favicon.ico")
}

func (h *handler) GetScript(ctx echo.Context) error {
	return ctx.File(h.rootPath + "/internal/view/graphPlayground/script.js")
}

func (h *handler) GetCSSLogin(ctx echo.Context) error {
	return ctx.File(h.rootPath + "/internal/view/graphPlayground/style.css")
}
