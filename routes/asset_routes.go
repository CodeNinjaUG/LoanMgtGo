package routes

import (
	"github.com/LoanMgtGo/controllers"
	"github.com/gin-gonic/gin"
)

func AssetRoutes(c *gin.Engine) {
	c.GET("/assets", controllers.GetAssets())
	c.GET("/asset/:asset_id", controllers.GetAsset())
	c.POST("/asset", controllers.CreateAsset())
	c.PATCH("/asset/:asset_id", controllers.UpdateAsset())
	c.DELETE("/asset/:asset_id", controllers.DeleteAsset())
}
