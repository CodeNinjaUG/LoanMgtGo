package routes

import (
	"github.com/LoanMgtGo/controllers"
	"github.com/gin-gonic/gin"
)

func JourneyEntryRoutes(c *gin.Engine) {
	c.GET("/journey/entries", controllers.GetJourneyEntries())
	c.GET("/journey/:entry_id", controllers.GetJourneyEntry())
	c.POST("/journey", controllers.CreateJourneyEntry())
	c.PATCH("/journey/entry/:entry_id", controllers.UpdateJourneyEntry())
	c.DELETE("/journey/entry/:entry_id", controllers.DeleteJourneyEntry())
}
