package models

import (
	"time"
)

type Author struct {
	ID			string		`json:"id" gorm:"primary_key;auto_increment"`
	Name		string		`json:"name"`
	Phone		string		`json:"phone"`
	Age			uint		`json:"age"`
	Status 		uint		`json:"status"`
	CreatedAt	time.Time	
	UpdatedAt	time.Time	
	DeletedAt	time.Time	
	
}