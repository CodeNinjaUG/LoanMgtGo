package routes

import (
	"github.com/LoanMgtGo/controllers"
	"github.com/gin-gonic/gin"
)

func  UserRoutes(c *gin.Engine) {
	c.GET("/users", controllers.GetUsers())
	c.GET("/users/:user_id", controllers.GetUser())
	c.POST("/users/signup", controllers.SignUp())
	c.POST("/users/login", controllers.Login())
}
