package controllers

import (
	"errors"
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

//getting loan products from the loanproduct model 
func (repo *LoanProductRepo) GetLoanProducts(c *gin.Context){
    var loanproducts []models.LoanProduct
    err:= models.GetLoanProducts(repo.Db, &loanproducts)
	if err!=nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error":err})
		return
	}
	c.JSON(http.StatusOK,loanproducts)
}

//getting a single loan from the database
func(repo *LoanProductRepo)GetLoanProduct(c *gin.Context){
       id,_ := c.Params.Get("id")
	   var loan_product models.LoanProduct
	   err:= models.GetLoanProduct(repo.Db,&loan_product,id)
	   if err!=nil{
		  if errors.Is(err , gorm.ErrRecordNotFound){
			  c.AbortWithStatus(http.StatusNotFound)
			  return
		  }
		    c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error":err})
		    return
	   }
	   c.JSON(http.StatusOK,loan_product)
}

//update loanproduct 
func (repo *LoanProductRepo)UpdateLoanProduct(c *gin.Context){
    var loan_product models.LoanProduct
	id,_:= c.Params.Get("id")
	err := models.GetLoanProduct(repo.Db,&loan_product,id)
	if err!=nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			c.AbortWithStatus(http.StatusNotFound)
			return 
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error":err})
		return
	}
	c.BindJSON(&loan_product)
	err= models.UpdateLoanProduct(repo.Db,&loan_product)
	if err!=nil{
		c.JSON(http.StatusOK,loan_product)
	}
}

//delete a loan product
func(repo *LoanProductRepo)DeleteLoanProduct(c *gin.Context){
   var loan_product models.LoanProduct
   id,_:= c.Params.Get("id")
   err:= models.DeleteLoanProduct(repo.Db,&loan_product,id)
   if err !=nil{
	   c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error":err})
	   return
   }
   c.JSON(http.StatusOK,gin.H{"message":"LoanProduct  Deleted Successfully"})
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
