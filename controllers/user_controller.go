package controllers

import (
	"net/http"

	"github.com/LoanMgtGo/database"
	"github.com/LoanMgtGo/helpers"
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

// type RegisterInput struct {
// 	Username string `json:"username" binding:"required"`
// 	Email    string `json:"email" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}


// func (repo *UserRepo) CreateUser(c *gin.Context) {
// 	var user models.User
// 	c.BindJSON(&user)
// 	err := models.CreateUser(repo.Db, &user)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
// 		return
// 	}
// 	c.JSON(http.StatusOK, user)
// }

func (repo *UserRepo) Login(c *gin.Context) {
	var user models.User
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	user.Name = input.Username
	user.Password = input.Password
	token, err := models.LoginCheck(repo.Db, user.Name, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, token)

}

func (repo *UserRepo) CurrentUser(c *gin.Context) {
	user_id, err := helpers.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := models.GetUserByID(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
func (repo *UserRepo) Register(c *gin.Context) {
	var registerinput RegisterInput
	var user models.User
	if err := c.ShouldBindJSON(&registerinput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Email = registerinput.Email
	user.Name = registerinput.Username
	user.Password = registerinput.Password
	_, err := models.CreateUser(repo.Db, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "registration successfully done"})
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
