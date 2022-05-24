package main

import (
	"net/http"

	"github.com/LoanMgtGo/controllers"
	"github.com/LoanMgtGo/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := setUpRouter()
	_ = r.Run(":8080")
}

func setUpRouter() *gin.Engine {

	r := gin.Default()
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	user_repo := controllers.NewUser()
	// public := r.Group("/api")
	r.POST("/user/register", user_repo.Register)
	r.POST("/user/login", user_repo.Login)
	protected := r.Group("/api")
	protected.Use(middleware.JwtAuthMiddleware())
	protected.GET("/user", user_repo.CurrentUser)
	loan_product_repo := controllers.NewLoanProduct()
	protected.GET("/loan/products", loan_product_repo.GetLoanProducts)
	protected.POST("/create/product", loan_product_repo.CreateLoanProduct)
	protected.GET("/loan/product/:id", loan_product_repo.GetLoanProduct)
	protected.PUT("/loan/product/:id", loan_product_repo.UpdateLoanProduct)
	protected.DELETE("/loan/product/:id", loan_product_repo.DeleteLoanProduct)

	loan_charge_repo := controllers.NewLoanCharge()
	protected.GET("/loan/charges", loan_charge_repo.GetLoanCharges)
	protected.POST("/create/charge", loan_charge_repo.CreateLoanCharge)
	protected.GET("/loan/charge/:id", loan_charge_repo.GetLoanCharge)
	protected.PUT("/loan/charge/:id", loan_charge_repo.UpdateLoanCharge)
	protected.DELETE("/loan/charge/:id", loan_charge_repo.DeleteLoanCharge)

	client_repo := controllers.NewClient()
	protected.GET("/clients", client_repo.GetClients)
	protected.POST("/create/client", client_repo.CreateClient)
	protected.GET("/client/:id", client_repo.GetClient)
	protected.PUT("/client/:id", client_repo.UpdateClient)
	protected.DELETE("/client/:id", client_repo.DeleteClient)

    chart_account_repo := controllers.NewChartAccount()
	protected.GET("/chartaccounts", chart_account_repo.GetChartsOfAccounts)
	protected.GET("/chartaccounts/:chart_account_id", chart_account_repo.GetChartAccount)
	protected.POST("/chartaccount", chart_account_repo.CreateChartAccount)
	protected.PATCH("/chartaccount/:chart_account_id", chart_account_repo.UpdateChartAccount)
	protected.DELETE("/chartaccount/:chart_account_id", chart_account_repo.DeleteChartAccount)

	return r
}
