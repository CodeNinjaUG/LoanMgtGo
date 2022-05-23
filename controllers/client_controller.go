package controllers

import (
	"errors"
	"net/http"

	"github.com/LoanMgtGo/database"
	"github.com/LoanMgtGo/models"
	"github.com/gin-gonic/gin"

	// "github.com/harranali/dump"
	"gorm.io/gorm"
)

type ClientRepo struct {
	Db *gorm.DB
}

func NewClient() *ClientRepo {
	db := database.InitDB()
	db.AutoMigrate(&models.Client{})
	return &ClientRepo{Db: db}
}

//creating a client method
func (repo *ClientRepo) CreateClient(c *gin.Context) {
	var client models.Client
	//dump.DD(c.BindJSON(&client))
	ir := c.BindJSON(&client)
	if ir != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ir})
		return
	}
	err := models.CreateClient(repo.Db, &client)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, client)
}

//getting a clients from the database
func (repo *ClientRepo) GetClients(c *gin.Context) {
	var clients []models.Client
	err := models.GetClients(repo.Db, &clients)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, clients)
}

//getting client by id
func (repo *ClientRepo) GetClient(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var client models.Client
	err := models.GetClient(repo.Db, &client, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, client)
}

//update a client
func (repo *ClientRepo) UpdateClient(c *gin.Context) {
	var client models.Client
	id, _ := c.Params.Get("id")
	err := models.GetClient(repo.Db, &client, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&client)
	err = models.UpdateClient(repo.Db, &client)
	if err != nil {
		c.JSON(http.StatusOK, client)
	}
}

//delete client
func (repo *ClientRepo) DeleteClient(c *gin.Context) {
	var client models.Client
	id, _ := c.Params.Get("id")
	err := models.DeleteClient(repo.Db, &client, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "LoanProduct  Deleted Successfully"})
}

// func GetClients() gin.HandlerFunc{
// 	return func(ctx *gin.Context) {

// 	}
// }

// func GetClient() gin.HandlerFunc{
// 	return func(ctx *gin.Context) {

// 	}
// }

// func CreateClient() gin.HandlerFunc{
// 	return func(ctx *gin.Context) {

// 	}
// }

// func UpdateClient() gin.HandlerFunc{
// 	return func(ctx *gin.Context) {

// 	}
// }
// func DeleteClient() gin.HandlerFunc{
// 	return func(ctx *gin.Context) {

// 	}
// }
