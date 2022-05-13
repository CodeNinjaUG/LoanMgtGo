package routes

import (
	"github.com/LoanMgtGo/controllers"
	"github.com/gin-gonic/gin"
)

func ShareRoutes(c *gin.Engine) {
	c.GET("/shares", controllers.GetShares())
	c.GET("/share/:share_id", controllers.GetShare())
	c.POST("/share", controllers.CreateShare())
	c.PATCH("/share/:share_id", controllers.UpdateShare())
	c.DELETE("/share/:share_id", controllers.DeleteShare())
}
