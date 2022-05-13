package main

import (
	"fmt"
	"os"

	"github.com/LoanMgtGo/database"
	"github.com/LoanMgtGo/routes"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	database.DB, err = gorm.Open("mysql", database.DB_url(database.InitDatabase()))
	if err != nil {
		fmt.Println("status", err)
	}
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	//router.Use(middleware.Authentication())
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
