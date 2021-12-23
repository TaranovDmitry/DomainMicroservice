package repository

import (
	"fmt"

	"github.com/TaranovDmitry/DomainMicroservice/entity"

	"github.com/jmoiron/sqlx"
)

type Ports struct {
	db *sqlx.DB
}

const (
	selectAllQuery = `SELECT id, name, isactive, company, email, phone, address, about, registered, latitude, longitude FROM ports`
	upsertQuery    = `INSERT INTO ports (id, name, isactive, company, email, phone, address, about, registered, latitude, longitude)
VALUES (:id, :name, :isactive, :company, :email, :phone, :address, :about, :registered, :latitude, :longitude)
ON CONFLICT (id) DO UPDATE SET
                               name = excluded.name,
                               isActive = excluded.isactive,
                               company = excluded.company,
                               email = excluded.email,
                               phone = excluded.phone,
                               address = excluded.address,
                               about = excluded.about,
                               registered = excluded.registered,
                               latitude = excluded.latitude,
                               longitude = excluded.longitude;`
)

func NewPortsRepository(db *sqlx.DB) *Ports {
	return &Ports{
		db: db,
	}
}

func (p Ports) FetchAllPortsFromDB() (entity.Ports, error) {
	var ports entity.Ports
	err := p.db.Select(&ports, selectAllQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to get all ports from DB: %w", err)
	}

	return ports, nil
}

func (p Ports) UpsertPortsInDB(ports entity.Ports) error {
	_, err := p.db.NamedExec(upsertQuery, ports)
	if err != nil {
		return fmt.Errorf("failed to upsert ports: %w", err)
	}
	return err
}
