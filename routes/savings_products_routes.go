package routes

import (
	"github.com/LoanMgtGo/controllers"
	"github.com/gin-gonic/gin"
)

func SavingsProductsRoutes(c *gin.Engine) {
	c.GET("/saving/products", controllers.GetSavingProducts())
	c.GET("/saving/product/:product_id", controllers.GetSavingProduct())
	c.POST("/saving/product", controllers.CreateSavingProduct())
	c.PATCH("/saving/product/:product_id", controllers.UpdateSavingProduct())
	c.DELETE("/saving/product/:saving_id", controllers.DeleteSavingProduct())
}
