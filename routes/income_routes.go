package routes

import (
	"github.com/LoanMgtGo/controllers"
	"github.com/gin-gonic/gin"
)

func IncomeRoutes(c *gin.Engine) {
	c.GET("/incomes", controllers.GetIncomes())
	c.GET("/income/:income_id", controllers.GetIncome())
	c.POST("/income", controllers.CreateIncome())
	c.PATCH("/expense/:expense_id", controllers.UpdateIncome())
	c.DELETE("/expense/:expense_id", controllers.DeleteIncome())
}
