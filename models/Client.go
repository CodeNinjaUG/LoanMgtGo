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
	Branch        string    `json:"branch"`
	ExternalID    string    `json:"external_id"`
	Title         Title     `json:"title"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Gender        string    `json:"gender"`
	MaritalStatus string    `json:"marital_status"`
	Mobile        int       `json:"mobile"`
	Country       string    `json:"country"`
	DOB           time.Time `json:"dob"`
	Staff         string    `json:"staff"`
	Email         string    `json:"email"`
	Profession    string    `json:"profession"`
	ClientType    string    `json:"client_type"`
	Photo         string    `json:"photo"`
	Address       string    `json:"address"`
	Notes         string    `json:"notes"`
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
func GetClients(db *gorm.DB , Clients *[]Client)(err error){
	err = db.Find(Clients).Error
	if err != nil{
		return err
	}
	return nil
}

//get client by id
func GetClient(db *gorm.DB, Client *Client, id string)(err error){
	err = db.Where("id=?",id).First(Client).Error
	if err!=nil{
		return err
	}
	return nil
}

//update client details
func UpdateClient(db *gorm.DB, Client *Client)(err error){
    db.Save(Client)
	return nil
}

func DeleteClient(db *gorm.DB, Client *Client, id string)(err error){
   db.Where("id=?",id).Delete(Client)
   return nil
}
