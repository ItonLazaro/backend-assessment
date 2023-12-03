package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//Define Users table for database communication

type Users struct {
	gorm.Model        //declares basic and defaults columns (ID and timestamps)
	Username   string `gorm:"size:255; not null, unique"`
	Password   string `gorm:"size:255; not null;"`
}

func (u *Users) BeforeSave() error {
	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}
