package service

import (
	"fmt"

	"github.com/zsandibe/effective_mobile_task/internal/domain"
	personRepository "github.com/zsandibe/effective_mobile_task/internal/repository/person"
)

type Person interface {
	FindPersonsList(searchParams domain.SearchParams) ([]domain.Person, error)
	FindPersonById(personId int) (domain.Person, error)
	AddPerson(person domain.Person) (domain.Person, error)
	DeletePersonById(id int) error
	UpdatePersonById(person domain.Person) error
}

type person struct {
	Enrichment
	repo personRepository.Person
}

func NewPerson(repo personRepository.Person, enrichment Enrichment) Person {
	return &person{
		Enrichment: enrichment,
		repo:       repo,
	}
}

func (p person) FindPersonsList(searchParams domain.SearchParams) ([]domain.Person, error) {
	return p.repo.FindPersonsBySearch(searchParams)
}

func (p person) FindPersonById(personId int) (domain.Person, error) {
	return p.repo.FindPersonById(personId)
}

func (p person) AddPerson(person domain.Person) (domain.Person, error) {
	age, err := p.GetPersonAgeByName(person.Name)
	if err != nil {
		return domain.Person{}, fmt.Errorf("failed to get age from api: %v", err)
	}
	person.Age = age

	gender, err := p.GetPersonGenderByName(person.Name)
	if err != nil {
		return domain.Person{}, fmt.Errorf("failed to get age from api: %v", err)
	}
	person.Gender = gender

	nationality, err := p.GetPersonNationalityByName(person.Name)
	if err != nil {
		return domain.Person{}, fmt.Errorf("failed to get nationality from api: %v", err)
	}
	person.Nationality = nationality

	id, err := p.repo.AddPerson(person)
	if err != nil {
		return domain.Person{}, fmt.Errorf("failed to add person: %v", err)
	}
	person.Id = id

	return person, nil
}

func (p person) DeletePersonById(id int) error {
	return p.repo.DeletePersonById(id)
}

func (p person) UpdatePersonById(person domain.Person) error {
	return p.repo.UpdatePersonById(person)
}
