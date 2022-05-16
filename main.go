package main

import (
	"net/http"
    "github.com/LoanMgtGo/controllers"
	"github.com/gin-gonic/gin"
)

var err error

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
	r.GET("/loan/products",loan_product_repo.GetLoanProducts)
	r.POST("/create/product",loan_product_repo.CreateLoanProduct)
	r.GET("/loan/product/:id",loan_product_repo.GetLoanProduct)
	r.PUT("/loan/product/:id",loan_product_repo.UpdateLoanProduct)
	r.DELETE("/loan/product/:id",loan_product_repo.DeleteLoanProduct)
	return r
}
