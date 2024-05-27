package handler

import (
	"github.com/labstack/echo/v4"
	"golang_graphs/backend/internal/handler/common"
)

func SetupRoutes(e *echo.Echo, com common.Handler) {
	e.POST("check_results", com.CheckResults)
	e.POST("auth_user", com.AuthUser)
	e.POST("create_user", com.CreateUser)
	e.POST("get_tasks_from_test", com.GetTasksFromTest)
	e.GET("get_tests", com.GetTests)
	e.POST("send_answers", com.SendAnswers)
	e.POST("insert_test", com.InsertTest)
	e.POST("insert_task", com.InsertTask)

	e.GET("ping", com.Ping)
}
