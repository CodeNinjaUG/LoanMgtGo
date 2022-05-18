package main

import (
	"net/http"
	"github.com/LoanMgtGo/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := setUpRouter()
	_ = r.Run(":8090")
}

func setUpRouter() *gin.Engine {

	r := gin.Default()
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

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
