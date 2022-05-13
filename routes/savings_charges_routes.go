package routes

import (
	"github.com/LoanMgtGo/controllers"
	"github.com/gin-gonic/gin"
)

func SavingsChargesRoutes(c *gin.Engine) {
	c.GET("/saving/charges", controllers.GetSavingCharges())
	c.GET("/saving/charge/:saving_id", controllers.GetSavingcharge())
	c.POST("/saving/charge", controllers.CreateSavingCharge())
	c.PATCH("/saving/charge/:saving_id", controllers.UpdateSavingCharge())
	c.DELETE("/saving/charge/:saving_id", controllers.DeleteSavingCharge())
}
