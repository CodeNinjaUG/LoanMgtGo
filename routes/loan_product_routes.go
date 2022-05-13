package routes

import (
	"github.com/LoanMgtGo/controllers"
	"github.com/gin-gonic/gin"
)

func LoanProductRoutes(c *gin.Engine) {
	c.GET("/loan/products", controllers.GetLoanProducts())
	c.GET("/loan/product/:product_id", controllers.GetLoanProduct())
	c.POST("/loan/product", controllers.CreateLoanProduct())
	c.PATCH("/loan/product/:product_id", controllers.UpdateLoanProduct())
	c.DELETE("/loan/product/:product_id", controllers.DeleteLoanProduct())
}
