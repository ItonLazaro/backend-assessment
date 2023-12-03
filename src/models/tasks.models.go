package models

import (
	"gorm.io/gorm"
)

//Define Tasks table for database communication

type Tasks struct {
	gorm.Model  //declares basic and defaults columns (ID and timestamps)
	Title       string
	Description string
	DueDate     string //should be changed to date data type
	Priority    string
	Status      string
}
