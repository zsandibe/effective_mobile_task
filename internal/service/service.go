package service

import (
	"github.com/zsandibe/effective_mobile_task/internal/repository"
	personService "github.com/zsandibe/effective_mobile_task/internal/service/person"
)

type Service struct {
	personService.Person
}

func NewPersonService(repo *repository.Repository) *Service {
	return &Service{
		Person: personService.NewPersonService(repo.Person),
	}
}
