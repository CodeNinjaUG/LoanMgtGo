package models

import (
	"time"

	"gorm.io/gorm"
)

type Title string

const (
	Mr   Title = "Mr."
	Mrs  Title = "Mrs."
	Dr   Title = "Dr"
	Miss Title = "Miss"
	Eng  Title = "Eng"
)

type Client struct {
	Branch        string    `json:"branch" binding:"required"`
	ExternalID    string    `json:"external_id" binding:"required"`
	Title         Title     `json:"title" binding:"required"`
	FirstName     string    `json:"first_name" binding:"required"`
	LastName      string    `json:"last_name" binding:"required"`
	Gender        string    `json:"gender" binding:"required"`
	MaritalStatus string    `json:"marital_status" binding:"required"`
	Mobile        string    `json:"mobile" binding:"required"`
	Country       string    `json:"country" binding:"required"`
	DOB           time.Time `json:"dob" binding:"required"`
	Staff         string    `json:"staff" binding:"required"`
	Email         string    `json:"email" binding:"required"`
	Profession    string    `json:"profession" binding:"required"`
	ClientType    string    `json:"client_type" binding:"required"`
	Photo         string    `json:"photo" binding:"required"`
	Address       string    `json:"address" binding:"required"`
	Notes         string    `json:"notes" binding:"required"`
	SubmittedOn   time.Time
}

//create client
func CreateClient(db *gorm.DB, Client *Client) (err error) {
	err = db.Create(Client).Error
	if err != nil {
		return err
	}
	return nil
}

//get clients
func GetClients(db *gorm.DB, Clients *[]Client) (err error) {
	err = db.Find(Clients).Error
	if err != nil {
		return err
	}
	return nil
}

//get client by id
func GetClient(db *gorm.DB, Client *Client, id string) (err error) {
	err = db.Where("id=?", id).First(Client).Error
	if err != nil {
		return err
	}
	return nil
}

//update client details
func UpdateClient(db *gorm.DB, Client *Client) (err error) {
	db.Save(Client)
	return nil
}

func DeleteClient(db *gorm.DB, Client *Client, id string) (err error) {
	db.Where("id=?", id).Delete(Client)
	return nil
}
