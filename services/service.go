package services

import (
	"fmt"

	"github.com/TaranovDmitry/Microservices/entity"
)

type PortRepository interface {
	FetchAllPortsFromDB() (entity.Ports, error)
	UpdatePortsInDB(ports entity.Ports) error
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
		return nil, fmt.Errorf("service failed to get datat from DB: %w", err)
	}

	return ports, err
}

func (s Service) Update(ports entity.Ports) error {
	return nil
}
