package repository

import (
	"github.com/jmoiron/sqlx"
	personRepository "github.com/zsandibe/effective_mobile_task/internal/repository/person"
)

type Repository struct {
	personRepository.Person
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Person: personRepository.NewPersonSql(db),
	}
}
