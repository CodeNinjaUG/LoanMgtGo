package routes

import (
	"github.com/LoanMgtGo/controllers"
	"github.com/gin-gonic/gin"
)

func ShareProductRoutes(c *gin.Engine) {
	c.GET("/share/products", controllers.GetShareProducts())
	c.GET("/share/product/:product_id", controllers.GetShareProduct())
	c.POST("/share/product", controllers.CreateShareProduct())
	c.PATCH("/share/product/:product_id", controllers.UpdateShareProduct())
	c.DELETE("/share/product/:product_id", controllers.DeleteShareProduct())
}
