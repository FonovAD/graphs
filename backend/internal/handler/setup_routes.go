package handler

import (
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, com Handler) {
	router := e.Group("")
	router.POST("create_user", com.CreateUser)
	router.POST("auth_user", com.AuthUser)

	authorizedRouter := e.Group("", com.AuthMiddleware())
	authorizedRouter.POST("check_results", com.CheckResults)
	authorizedRouter.POST("get_tasks_from_test", com.GetTasksFromTest)
	authorizedRouter.GET("get_tests", com.GetTests)
	authorizedRouter.POST("send_answers", com.SendAnswers)
	authorizedRouter.POST("insert_test", com.InsertTest)
	authorizedRouter.POST("insert_task", com.InsertTask)
}
