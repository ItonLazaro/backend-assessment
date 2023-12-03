package routes

import (
	"example/todo-go/src/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	route := gin.Default()

	route.POST("/api/login", controllers.Login)
	// route.POST("/api/logout", controllers.logout)
	route.POST("/api/register", controllers.Register)

	route.GET("/api/task", controllers.GetAllTasks)
	route.POST("/api/task", controllers.CreateTask)
	route.PUT("/api/task/:id", controllers.UpdateTask)
	route.DELETE("/api/task/:id", controllers.DeleteTask)

	//Run a Route when triggered

	route.Run("localhost:8080")
}
