package routes

import (
	"example/todo-go/src/controllers"
	"example/todo-go/src/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes() {
	route := gin.Default()

	route.POST("/api/login", controllers.Login)
	// route.POST("/api/logout", controllers.logout)
	route.POST("/api/register", controllers.Register)

	route.Use(middlewares.JwtAuthMiddleware()).GET("/api/task", controllers.GetAllTasks)
	route.Use(middlewares.JwtAuthMiddleware()).POST("/api/task", controllers.CreateTask)
	route.Use(middlewares.JwtAuthMiddleware()).PUT("/api/task/:id", controllers.UpdateTask)
	route.Use(middlewares.JwtAuthMiddleware()).DELETE("/api/task/:id", controllers.DeleteTask)

	//Run a Route when triggered

	route.Run("localhost:8080")
}
