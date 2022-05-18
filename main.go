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
	r.GET("/user/register", user_repo.Register)
	r.GET("/user/login", user_repo.Login)
	protected := r.Group("/api/admin")
	protected.Use(middleware.JwtAuthMiddleware())
	protected.GET("/user", user_repo.CurrentUser)
	loan_product_repo := controllers.NewLoanProduct()
	r.GET("/loan/products", loan_product_repo.GetLoanProducts)
	r.POST("/create/product", loan_product_repo.CreateLoanProduct)
	r.GET("/loan/product/:id", loan_product_repo.GetLoanProduct)
	r.PUT("/loan/product/:id", loan_product_repo.UpdateLoanProduct)
	r.DELETE("/loan/product/:id", loan_product_repo.DeleteLoanProduct)

	loan_charge_repo := controllers.NewLoanCharge()
	r.GET("/loan/charges", loan_charge_repo.GetLoanCharges)
	r.POST("/create/charge", loan_charge_repo.CreateLoanCharge)
	r.GET("/loan/charge/:id", loan_charge_repo.GetLoanCharge)
	r.PUT("/loan/charge/:id", loan_charge_repo.UpdateLoanCharge)
	r.DELETE("/loan/charge/:id", loan_charge_repo.DeleteLoanCharge)

	return r
}
