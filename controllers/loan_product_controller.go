package controllers

import (
	"net/http"

	"github.com/LoanMgtGo/database"
	"github.com/LoanMgtGo/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoanProductRepo struct {
	Db *gorm.DB
}

func New() *LoanProductRepo {
	db := database.InitDB()
	db.AutoMigrate(&models.LoanProduct{})
	return &LoanProductRepo{Db: db}
}

//creating loan product reffering a model method createloanproduct
func (repo *LoanProductRepo) CreateLoanProduct(c *gin.Context) {
	var loanproduct models.LoanProduct
	c.BindJSON(&loanproduct)
	err := models.CreateLoanProduct(repo.Db, &loanproduct)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, loanproduct)
}

// func GetLoanProducts() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {

// 	}
// }
// func GetLoanProduct() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {

// 	}
// }
// func CreateLoanProduct() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {

// 	}
// }
// func UpdateLoanProduct() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {

// 	}
// }
// func DeleteLoanProduct() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {

// 	}
// }
