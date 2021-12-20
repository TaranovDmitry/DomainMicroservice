package service

import "github.com/TaranovDmitry/Microservices/repository"

type TodoList interface {

	
}

type Service struct {
	TodoList
}

func NewService(repos *repository.Repository) *Service  {
	return &Service{}
}