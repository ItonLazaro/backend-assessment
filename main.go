package main

import (
	"example/todo-go/src/config"
	"example/todo-go/src/routes"

	"gorm.io/gorm"
)

var (
	//set a global variable "db"
	db *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.DisconnectDB(db)

	//load all the routes declared in /src/routes file
	routes.Routes()
}
