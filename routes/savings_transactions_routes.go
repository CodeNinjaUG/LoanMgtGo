package routes

import (
	"github.com/LoanMgtGo/controllers"
	"github.com/gin-gonic/gin"
)

func SavingsTransactionsRoutes(c *gin.Engine) {
	c.GET("/saving/transactions", controllers.GetSavingTransactions())
	c.GET("/saving/transaction/:transaction_id", controllers.GetSavingTransaction())
	c.POST("/saving/transaction", controllers.CreateSavingTransaction())
	c.PATCH("/saving/product/:transaction_id", controllers.UpdateSavingTransaction())
	c.DELETE("/saving/product/:transaction_id", controllers.DeleteSavingTransaction())
}
