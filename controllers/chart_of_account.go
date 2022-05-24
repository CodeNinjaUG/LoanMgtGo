package controllers

import (
	"errors"
	"net/http"

	"github.com/LoanMgtGo/database"
	"github.com/LoanMgtGo/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ChartAccountRepo struct{
	Db *gorm.DB
}

func NewChartAccount() *ChartAccountRepo{
   db := database.InitDB()
   db.AutoMigrate(&models.ChartAccount{})
   return &ChartAccountRepo{Db:db}
}

//creating a chartaccount 
func (repo *ChartAccountRepo) CreateChartAccount(c *gin.Context){
	var chart_account models.ChartAccount
	c.BindJSON(&chart_account)
    err := models.CreateAccount(repo.Db,&chart_account) 
	if err!=nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error":err})
		return 
	}
	c.JSON(http.StatusOK,chart_account)
}

//getting chartaccounts 
func(repo *ChartAccountRepo) GetChartsOfAccounts(c *gin.Context){
	var chart_accounts [] models.ChartAccount
	err := models.GetAccounts(repo.Db,&chart_accounts)
	if err!=nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error":err})
		return 
	}
	c.JSON(http.StatusOK,chart_accounts)
}

//getting a chartaccount
func(repo *ChartAccountRepo) GetChartAccount(c *gin.Context){
	id,_ := c.Params.Get("id")
	var chart_account models.ChartAccount
	err := models.GetAccount(repo.Db,&chart_account, id)
	if err!=nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error":err})
		return
	}
	c.JSON(http.StatusOK,chart_account)
}

//updating a chartaccount
func(repo *ChartAccountRepo) UpdateChartAccount(c *gin.Context){
	id,_ := c.Params.Get("id")
	var chart_account models.ChartAccount
	err:= models.GetAccount(repo.Db, &chart_account, id)
	if errors.Is(err, gorm.ErrRecordNotFound){
		c.AbortWithStatus(http.StatusNotFound)
		return 
	}
	c.BindJSON(&chart_account)
	err = models.UpdateAccount(repo.Db,&chart_account)
	if err!=nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error": err})
		return 
	}
	c.JSON(http.StatusOK, chart_account)
}
//deleting a chart account 
func (repo *ChartAccountRepo) DeleteChartAccount(c *gin.Context){
	id,_ := c.Params.Get("id")
	var chart_account models.ChartAccount
	err:= models.DeleteAccount(repo.Db,&chart_account,id)
	if err!=nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error":err})
		return
	}
	c.JSON(http.StatusOK,chart_account)
}

// func GetChartsOfAccounts() gin.HandlerFunc{
// 	return func(ctx *gin.Context){

// 	}
// }

// func CreateChartAccount() gin.HandlerFunc{
// 	return func(ctx *gin.Context) {

// 	}
// }

// func GetChartAccount() gin.HandlerFunc{
// 	return func(ctx *gin.Context) {

// 	}
// }

// func UpdateChartAccount() gin.HandlerFunc{
// 	return func(ctx *gin.Context) {}
// }


// func DeleteChartAccount() gin.HandlerFunc{
// 	return func(ctx *gin.Context) {}
// }
