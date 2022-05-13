package main

import (
	"os"

	"github.com/LoanMgtGo/middleware"
	"github.com/LoanMgtGo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.AssetRoutes(router)
	routes.AssetTypeRoutes(router)
	routes.ChartsAccountsRoutes(router)
	routes.ClientRoutes(router)
	routes.ExpensesRoutes(router)
	routes.IncomeRoutes(router)
	routes.JourneyEntryRoutes(router)
	routes.LoanRoutes(router)
	routes.LoanChargeRoutes(router)
	routes.LoanProductRoutes(router)
	routes.SavingsRoutes(router)
	routes.SavingsChargesRoutes(router)
	routes.SavingsProductsRoutes(router)
	routes.SavingsTransactionsRoutes(router)
	routes.ChartsAccountsRoutes(router)
	routes.ShareRoutes(router)
	routes.ShareChargesRoutes(router)
	routes.ShareProductRoutes(router)
}
