package routes

import (
	"github.com/LoanMgtGo/controllers"
	"github.com/gin-gonic/gin"
)

func LoanRoutes(c *gin.Engine) {
	c.GET("/loans", controllers.GetLoans())
	c.GET("/loan/:loan_id", controllers.GetLoan())
	c.POST("/loan", controllers.CreateLoan())
	c.PATCH("/loan/:loan_id", controllers.UpdateLoan())
	c.DELETE("/loan/:loan_id", controllers.DeleteLoan())
}
