package controllers

import (
	"net/http"

	"github.com/LoanMgtGo/database"
	"github.com/LoanMgtGo/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func NewUser() *UserRepo {
	db := database.InitDB()
	db.AutoMigrate(&models.User{})
	return &UserRepo{Db: db}
}

// type UserRepo struct {
// 	Username string `json:"username" binding:"required"`
// 	Email    string `json:"email" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }

func (repo *UserRepo) CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := models.CreateUser(repo.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

// func GetUsers() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 	}
// }

// func GetUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 	}
// }

// func SignUp() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 	}
// }

// func Login() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {

// 	}
// }

// func HashPassword(password string) string {

// }

// func VerifyPassword(userPassword string,providedPassword string)(bool,string){

// }
