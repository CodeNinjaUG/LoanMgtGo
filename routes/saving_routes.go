package routes

import (
	"github.com/LoanMgtGo/controllers"
	"github.com/gin-gonic/gin"
)

func SavingsRoutes(c *gin.Engine) {
	c.GET("/savings", controllers.GetSavings())
	c.GET("/saving/:saving_id", controllers.GetSaving())
	c.POST("/saving", controllers.CreateSaving())
	c.PATCH("/saving/:saving_id", controllers.UpdateSaving())
	c.DELETE("/saving/:saving_id", controllers.DeleteSaving())
}
