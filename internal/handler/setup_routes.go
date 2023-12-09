package handler

import (
	"github.com/labstack/echo/v4"
	"golang_graphs/internal/handler/common"
	"golang_graphs/internal/handler/files"
	"golang_graphs/internal/handler/task"
)

func SetupRoutes(e *echo.Echo, com common.Handler, files files.Handler, task task.Handler) {
	e.GET("check_results", com.CheckResults)
	e.POST("auth_user", com.AuthUser)
	e.POST("create_user", com.CreateUser)
	e.POST("get_tasks_from_test", com.GetTasksFromTest)
	e.GET("get_tests", com.GetTests)

	e.POST("task_components", task.TaskComponents)
	e.POST("task_is_euler", task.TaskIsEulerUndirected)
	e.POST("task_is_bipartition", task.TaskIsBipartition)

	e.GET("css/style.css", files.GetCSS)
	e.GET("js/d3.v5.min.js", files.GetJSD3)
	e.GET("js/script.js", files.GetJS)
	e.GET("favicon.ico", files.Favicon)
	e.GET("script", files.GetScript)
	e.GET("css", files.GetCSSLogin)
}
