package controllers

import (
	"example/todo-go/src/config"
	"example/todo-go/src/models"
	token "example/todo-go/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var dbUsers *gorm.DB = config.ConnectDB()

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	UserRequest
	ID uint `json:"id"`
}

func Register(context *gin.Context) {
	var data UserRequest

	isValidRequest := context.ShouldBindJSON(&data)

	//Throw a HTTP Error 400 Bad Request
	if isValidRequest != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": isValidRequest.Error()})
		return
	}

	//initialize User Model
	user := models.Users{}
	user.Username = data.Username
	hashedPassword, err := hashPassword(data.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong with registering a new user"})
		return
	}
	user.Password = hashedPassword

	result := dbUsers.Create(&user)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong with registering a new user"})
		return
	}

	//build a response data
	var response UserResponse
	response.ID = user.ID
	response.Username = user.Username

	//create a http OK response
	context.JSON(http.StatusCreated, response)
}

func Login(context *gin.Context) {
	var data UserRequest

	isValidRequest := context.ShouldBindJSON(&data)

	//Throw a HTTP Error 400 Bad Request
	if isValidRequest != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": isValidRequest.Error()})
		return
	}

	user := models.Users{}
	user.Username = data.Username
	user.Password = data.Password

	findUsername := dbUsers.Where("username = ?", user.Username).First(&user)

	if findUsername.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User not existing"})
		return
	}

	err := verifyPassword(data.Password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Login Error"})
		return
	}

	token, err := token.GenerateToken(user.ID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Login Error"})
		return
	}

	context.JSON(http.StatusOK, token)
}

// methods that starts with lowercase are private methods
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
