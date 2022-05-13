package routes

import (
	"github.com/LoanMgtGo/controllers"
	"github.com/gin-gonic/gin"
)

func ShareChargesRoutes(c *gin.Engine) {
	c.GET("/share/charges", controllers.GetShareCharges())
	c.GET("/share/charge/:charge_id", controllers.GetShareCharge())
	c.POST("/share/charge", controllers.CreateShareCharge())
	c.PATCH("/share/charge/:charge_id", controllers.UpdateShareCharge())
	c.DELETE("/share/charge/:charge_id", controllers.DeleteShareCharge())
}
