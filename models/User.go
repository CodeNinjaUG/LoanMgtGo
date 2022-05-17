package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"uniqueIndex; type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;notnull" json:"-"`
}

func CreateUser(db *gorm.DB, User *User) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

func BeforeSave(db *gorm.DB, User *User) (err error) {
	//turns password to hash
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(User.Password),bcrypt.DefaultCost)
	if err!=nil{
		return err
	}
	User.Password = string(hashedpassword)
	//remove spaces in the password
	User.Password = html.EscapeString(strings.TrimSpace(User.Password))
	return nil
}
