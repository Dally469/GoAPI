package models

import (
	"time"
)



type Book struct{
	Id          	string 		`json:"id" gorm:"primary_key;auto_increment" default:1`
	Name 			string 		`json:"name"`
	AuthorId 		string 		`json:"author_id"`
	Publication 	string 		`json:"publication"`
	Status       	uint32 		`json:"status"`
	CreatedAt    	time.Time	
	UpdatedAt    	time.Time
}
