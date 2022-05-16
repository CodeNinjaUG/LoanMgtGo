package controllers

import (
	"errors"
	"net/http"

	"github.com/LoanMgtGo/database"
	"github.com/LoanMgtGo/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoanChargeRepo struct {
	Db *gorm.DB
}

func NewLoanCharge() *LoanChargeRepo {
	db := database.InitDB()
	db.AutoMigrate(&models.LoanCharge{})
	return &LoanChargeRepo{Db: db}
}

//creating loan product reffering a model method createloanproduct
func (repo *LoanChargeRepo) CreateLoanCharge(c *gin.Context) {
	var loan_charge models.LoanCharge
	c.BindJSON(&loan_charge)
	err := models.CreateLoanCharge(repo.Db, &loan_charge)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, loan_charge)
}

//getting loan products from the loanproduct model
func (repo *LoanChargeRepo) GetLoanCharges(c *gin.Context) {
	var loan_charge []models.LoanCharge
	err := models.GetLoanCharges(repo.Db, &loan_charge)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, loan_charge)
}

//getting a single loan from the database
func (repo *LoanChargeRepo) GetLoanCharge(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var loan_charge models.LoanCharge
	err := models.GetLoanCharge(repo.Db, &loan_charge, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, loan_charge)
}

//update loanproduct
func (repo *LoanChargeRepo) UpdateLoanCharge(c *gin.Context) {
	var loan_charge models.LoanCharge
	id, _ := c.Params.Get("id")
	err := models.GetLoanCharge(repo.Db, &loan_charge, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&loan_charge)
	err = models.UpdateLoanCharge(repo.Db, &loan_charge)
	if err != nil {
		c.JSON(http.StatusOK, loan_charge)
	}
}

//delete a loan product
func (repo *LoanChargeRepo) DeleteLoanCharge(c *gin.Context) {
	var loan_charge models.LoanCharge
	id, _ := c.Params.Get("id")
	err := models.DeleteLoanCharge(repo.Db, &loan_charge, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "LoanCharge  Deleted Successfully"})
}

// func GetLoanCharges() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {

// 	}
// }
// func GetLoanCharge() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {

// 	}
// }
// func CreateLoanCharge() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {

// 	}
// }
// func UpdateLoanCharge() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {

// 	}
// }
// func DeleteLoanCharge() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {

// 	}
// }
