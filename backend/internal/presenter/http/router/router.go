package router

import (
	"golang_graphs/backend/internal/presenter/http/handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.GET("/ping", Ping)

	userRouter := e.Group("/api/v1/user")
	userRouter.POST("/auth_user", h.AuthUser)
	userRouter.POST("/create_user", h.CreateUser)

	teacherRouter := e.Group("/api/v1/teacher", h.TeacherMiddleware())
	teacherRouter.GET("/modules", h.GetModules)
	teacherRouter.POST("/create_lab", h.CreateLab)
	teacherRouter.POST("/lab_info", h.GetLabInfo)
	teacherRouter.POST("/remove_user_lab", h.RemoveUserLab)
	teacherRouter.PATCH("/update_lab_info", h.UpdateLabInfo)
	teacherRouter.POST("/assigne_lab", h.AssignLab)
	teacherRouter.POST("/assigne_lab_group", h.AssignLabGroup)
	teacherRouter.POST("/module_lab", h.AddModuleLab)
	teacherRouter.POST("/remove_module_lab", h.RemoveModuleLab)
	teacherRouter.POST("/not_assigned_labs", h.GetNonAssignedLabs)
	teacherRouter.GET("/assigned_labs", h.GetAssignedLabs)
	teacherRouter.POST("/lab_modules", h.GetLabModules)
	teacherRouter.GET("/groups", h.GetGroups)
	teacherRouter.POST("/create_task", h.CreateTask)
	teacherRouter.PATCH("/update_task", h.UpdateTask)
	teacherRouter.POST("/tasks_by_module", h.GetTasksByModule)
	teacherRouter.POST("/results", h.GetLabResults)
	teacherRouter.POST("/students", h.GetStudents)

	studentRouter := e.Group("/api/v1/student", h.StudentMiddleware())
	studentRouter.POST("/assigned_tasks_by_module", h.GetAssignedTasksByModule)
	studentRouter.POST("/begin_lab", h.BeginLab)
	studentRouter.POST("/finish_lab", h.FinishLab)
}

func Ping(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "pong")
}
