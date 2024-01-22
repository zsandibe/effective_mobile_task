package service

import (
	"github.com/zsandibe/effective_mobile_task/config"
	"github.com/zsandibe/effective_mobile_task/internal/repository"
	personService "github.com/zsandibe/effective_mobile_task/internal/service/person"
	personApiService "github.com/zsandibe/effective_mobile_task/internal/service/personApi"
)

type Service struct {
	personService.Person
	personApiService.PersonApi
}

func NewPersonService(repo *repository.Repository, config *config.Config) *Service {
	return &Service{
		Person:    personService.NewPersonService(repo.Person),
		PersonApi: personApiService.NewPersonApiService(config),
	}
}
