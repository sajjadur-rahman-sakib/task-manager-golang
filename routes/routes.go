package routes

import (
	"task-manager-golang/controllers"
	"task-manager-golang/middleware"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.POST("/signup", controllers.Signup)
	e.POST("/login", controllers.Login)

	auth := e.Group("/tasks", middleware.JWTMiddleware)
	auth.POST("", controllers.CreateTask)
	auth.GET("", controllers.GetTasks)
	auth.PUT("/:id", controllers.UpdateTask)
	auth.DELETE("/:id", controllers.DeleteTask)
	auth.PUT("/:id/complete", controllers.CompleteTask)
}
