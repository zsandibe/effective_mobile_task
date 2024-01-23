package service

import (
	"github.com/zsandibe/effective_mobile_task/config"
	"github.com/zsandibe/effective_mobile_task/internal/repository"
)

type Service struct {
	Person
}

func NewService(repo *repository.Repository, config *config.Config) *Service {
	enrichment := NewEnrichment(config)
	return &Service{
		Person: NewPerson(repo.Person, enrichment),
	}
}
