package services

import (
	"fmt"
	"github.com/TaranovDmitry/Microservices/entity"
)

type PortRepository interface {
	AllPorts() (entity.Ports, error)
	UpdatePort(ports entity.Ports) error
}

type Service struct {
	portRepo PortRepository
}

func NewService(pr PortRepository) *Service {
	return &Service{
		portRepo: pr,
	}
}

func (s Service) FetchAllPortsFromDB() (entity.Ports, error) {
	ports, err := s.portRepo.AllPorts()
	if err != nil {
		return nil, fmt.Errorf("service failed to get datat from DB: %w", err)
	}

	return ports, err
}
