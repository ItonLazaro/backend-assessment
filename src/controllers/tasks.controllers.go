package controllers

import (
	"example/todo-go/src/config"
	"example/todo-go/src/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

var db *gorm.DB = config.ConnectDB() //rename db

// Task structure for request body
type taskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Priority    string `json:"priority"`
	Status      string `json:"status"`
}

// Declare request body for Task
type taskReponse struct {
	taskRequest
	ID uint `json:"id"`
}

func CreateTask(context *gin.Context) {
	var data taskRequest

	isValidRequest := context.ShouldBindJSON(&data)

	//Throw a HTTP Error 400 Bad Request
	if isValidRequest != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": isValidRequest.Error()})
		return
	}

	task := models.Tasks{}
	task.Title = data.Title
	task.Description = data.Description
	task.DueDate = data.DueDate
	task.Priority = data.Priority
	task.Status = data.Status

	result := db.Create(&task)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	//build a response data
	var response taskReponse
	response.ID = task.ID
	response.Title = task.Title
	response.Description = task.Description
	response.DueDate = task.DueDate
	response.Priority = task.Priority
	response.Status = task.Status

	//Create a http OK response
	context.JSON(http.StatusCreated, response)
}

func GetAllTasks(context *gin.Context) {
	var tasks []models.Tasks

	err := db.Find(&tasks)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting all tasks"})
		// context.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	// if result.RowsAffected == 0 {
	// 	return Response{
	// 		Status:  http.StatusNotFound,
	// 		Message: "No records found",
	// 	}, nil
	// }

	//Create a http response
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    tasks,
	})
}

func UpdateTask(context *gin.Context) {
	var data taskRequest

	//Defining request parameter to get task id
	idTask := cast.ToUint(context.Param("id"))
	// idTask := cast.ToUint(reqParamId)

	isValidRequest := context.BindJSON(&data) //used twice. Can be moved to a separate function (TODO)

	//Throw a HTTP Error 400 Bad Request
	if isValidRequest != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": isValidRequest.Error()})
		return
	}

	//Initiate the Task model
	task := models.Tasks{}

	//Query to find the task by id from the parameter
	findTask := db.Where("id = ?", idTask).First(&task)

	if findTask.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Cannot find task"})
		// context.JSON(http.StatusBadRequest, gin.H{"error": findTask.Error}) //This should be logged for debugging purposes (TODO)
		return
	}

	task.Title = data.Title
	task.Description = data.Description
	task.DueDate = data.DueDate
	task.Priority = data.Priority
	task.Status = data.Status

	fmt.Println(task)

	result := db.Save(&task)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error updating the task"})
		// context.JSON(http.StatusBadRequest, result)
		return
	}

	// create a response body
	var response taskReponse
	response.ID = task.ID
	response.Title = task.Title
	response.Description = task.Description
	response.DueDate = task.DueDate
	response.Priority = task.Priority
	response.Status = task.Status

	//create a http response
	context.JSON(http.StatusCreated, response)
}

func DeleteTask(context *gin.Context) {
	task := models.Tasks{}

	//Defining request parameter to get task id
	idTask := context.Param("id")
	// idTask := cast.ToUint(reqParamId)

	db.Where("id = ?", idTask).Delete(&task)

	//creating a http response
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    idTask,
	})
}
