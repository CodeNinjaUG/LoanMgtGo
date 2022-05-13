package routes

import (
	"github.com/LoanMgtGo/controllers"
	"github.com/gin-gonic/gin"
)

func ExpensesRoutes(c *gin.Engine) {
	c.GET("/expenses", controllers.GetExpenses())
	c.GET("/expense/:expense_id", controllers.GetExpense())
	c.POST("/expense", controllers.CreateExpense())
	c.PATCH("/expense/:expense_id", controllers.UpdateExpense())
	c.DELETE("/expense/:expense_id", controllers.DeleteExpense())
}
