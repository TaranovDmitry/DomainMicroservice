package services

import (
	"fmt"

	"github.com/TaranovDmitry/DomainMicroservice/entity"
)

type PortRepository interface {
	FetchAllPortsFromDB() (entity.Ports, error)
	UpsertPortsInDB(ports entity.Ports) error
}

type Service struct {
	portRepo PortRepository
}

func NewService(pr PortRepository) *Service {
	return &Service{
		portRepo: pr,
	}
}

func (s Service) AllPorts() (entity.Ports, error) {
	ports, err := s.portRepo.FetchAllPortsFromDB()
	if err != nil {
		return nil, fmt.Errorf("service failed to get data from DB: %w", err)
	}

	return ports, err
}

func (s Service) Upsert(ports entity.Ports) error {
	err := s.portRepo.UpsertPortsInDB(ports)
	if err != nil {
		return fmt.Errorf("service failed to upsert data in DB: %w", err)
	}
	return nil
}
