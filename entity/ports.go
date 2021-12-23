package entity

import "time"

type Port struct {
	ID         int       `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	IsActive   bool      `json:"isActive" db:"isactive"`
	Company    string    `json:"company" db:"company"`
	Email      string    `json:"email" db:"email"`
	Phone      string    `json:"phone" db:"phone"`
	Address    string    `json:"address" db:"address"`
	About      string    `json:"about" db:"about"`
	Registered time.Time `json:"registered" db:"registered"`
	Latitude   float64   `json:"latitude" db:"latitude"`
	Longitude  float64   `json:"longitude" db:"longitude"`
}

type Ports []Port
