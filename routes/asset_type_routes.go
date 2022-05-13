package routes

import (
	"github.com/LoanMgtGo/controllers"
	"github.com/gin-gonic/gin"
)

func AssetTypeRoutes(c *gin.Engine) {
	c.GET("/assettypes", controllers.GetAssetTypes())
	c.GET("/assettype/:asset_type_id", controllers.GetAssetType())
	c.POST("/assettype", controllers.CreateAssetType())
	c.PATCH("/asset/:asset_id", controllers.UpdateAssetType())
	c.DELETE("/asset/:asset_id", controllers.DeleteAssetType())
}
