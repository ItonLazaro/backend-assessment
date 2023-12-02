package routes

import (
	"example/todo-go/src/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	route := gin.Default()

	route.GET("/task", controllers.GetAllTasks)
	route.POST("/task", controllers.CreateTask)
	route.PUT("/task/:id", controllers.UpdateTask)
	route.DELETE("/task/:id", controllers.DeleteTask)

	//Run a Route when triggered

	route.Run("localhost:8080")
}
