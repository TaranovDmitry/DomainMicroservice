package repository

import (
	"fmt"

	"github.com/TaranovDmitry/Microservices/entity"

	"github.com/jmoiron/sqlx"
)

type Ports struct {
	db *sqlx.DB
}

func NewPortsRepository(db *sqlx.DB) *Ports {
	return &Ports{
		db: db,
	}
}

func (p Ports) FetchAllPortsFromDB() (entity.Ports, error) {
	var ports entity.Ports
	err := p.db.Select(&ports, "SELECT * FROM ports")
	if err != nil {
		return nil, fmt.Errorf("failed to get all ports from DB: %w", err)
	}

	return ports, nil
}

func (p Ports) UpdatePortsInDB(ports entity.Ports) error {
	return nil
}
