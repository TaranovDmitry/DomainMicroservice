package entity

import "time"

type Port struct {
	ID         int
	Name       string
	IsActive   bool
	Company    string
	Email      string
	Phone      string
	Address    string
	About      string
	Registered time.Time
	Latitude   float64
	Longitude  float64
}

type Ports []Port
