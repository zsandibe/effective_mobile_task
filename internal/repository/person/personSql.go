package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/zsandibe/effective_mobile_task/internal/domain"
)

type PersonSql struct {
	db *sqlx.DB
}
type Person interface {
	FindPersonsBySearch(searchParams domain.SearchParams) ([]domain.Person, error)
	FindPersonById(personId int) (domain.Person, error)
	AddPerson(person domain.Person) (int, error)
	DeletePersonById(id int) error
	UpdatePersonById(person domain.Person) error
}

func NewPersonSql(db *sqlx.DB) *PersonSql {
	return &PersonSql{
		db: db,
	}
}
