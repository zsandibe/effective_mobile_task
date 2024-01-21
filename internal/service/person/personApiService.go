package person

import (
	"github.com/zsandibe/effective_mobile_task/internal/domain"
	personRepository "github.com/zsandibe/effective_mobile_task/internal/repository/person"
)

type Person interface {
	FindPersonsList(searchParams domain.SearchParams) ([]domain.Person, error)
	FindPersonById(personId int) (domain.Person, error)
	AddPerson(person domain.Person) (int, error)
	DeletePersonById(personId int) error
	UpdatePersonById(person domain.Person) error
}

type PersonService struct {
	repo personRepository.Person
}

func NewPersonService(repo personRepository.Person) *PersonService {
	return &PersonService{
		repo: repo,
	}
}
