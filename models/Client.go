package models

import (
	"time"
)

type Title string

const (
	Mr      Title =     "Mr."
	Mrs     Title =     "Mrs."
	Dr      Title =     "Dr"
	Miss    Title =     "Miss"
	Eng     Title =     "Eng"
)


type Client struct{
	Branch string            `json:"branch"`
	ExternalID string        `json:"external_id"`
	Title Title              `json:"title"`
	FirstName  string        `json:"first_name"`
	LastName   string        `json:"last_name"`
	Gender     string        `json:"gender"`
	MaritalStatus string     `json:"marital_status"`
	Mobile       int         `json:"mobile"`
	Country      string      `json:"country"`
    DOB         time.Time    `json:"dob"`
	Staff       string       `json:"staff"`
	Email       string       `json:"email"`
	Profession  string       `json:"profession"`
	ClientType  string       `json:"client_type"`
	Photo       string       `json:"photo"`
	Address     string       `json:"address"`
	Notes       string       `json:"notes"`
	SubmittedOn time.Time    
	}

