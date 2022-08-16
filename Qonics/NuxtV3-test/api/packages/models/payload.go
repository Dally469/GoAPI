package models

import "time"

type Payload struct {
	ID         string
	Name       string
	EntityId   string
	Email      string
	Phone      string
	EntityName string
	Time       time.Time
}
