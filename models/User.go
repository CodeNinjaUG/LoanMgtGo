package models

import (
	"html"
	"strings"

	"errors"

	"github.com/LoanMgtGo/helpers"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primary_key:auto_increment" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"uniqueIndex; type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;notnull" json:"-"`
}

func CreateUser(db *gorm.DB, u *User) (*User, error) {
	var err error
	err = db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func BeforeSave(db *gorm.DB, User *User) (err error) {
	//turns password to hash
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	User.Password = string(hashedpassword)
	//remove spaces in the password
	User.Password = html.EscapeString(strings.TrimSpace(User.Password))
	return nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(db *gorm.DB, username string, password string) (string, error) {
	var err error
	u := User{}
	err = db.Model(User{}).Where("name=?", username).Take(&u).Error
	if err != nil {
		return "", err
	}
	err = VerifyPassword(password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	//n, err := int(u.ID)
	token, err := helpers.GenerateToken(u.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
func GetUserByID(uid uint,db *gorm.DB) (User, error) {
	var u User
	if err := db.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found")
	}
	u.PrepareGive()
	return u, nil
}

func (u *User) PrepareGive() {
	u.Password = ""
}
