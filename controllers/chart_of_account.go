package controllers

import (
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
