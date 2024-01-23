package person

import (
	"fmt"

	"github.com/zsandibe/effective_mobile_task/internal/domain"
)

func (r *PersonService) FindPersonsList(searchParams domain.SearchParams) ([]domain.Person, error) {
	return r.repo.FindPersonsBySearch(searchParams)
}

func (r *PersonService) FindPersonById(personId int) (domain.Person, error) {
	return r.repo.FindPersonById(personId)
}

func (r *PersonService) AddPerson(person domain.Person) (domain.Person, error) {
	age, err := r.api.GetPersonAgeByName(person.Name)
	if err != nil {
		return domain.Person{}, fmt.Errorf("failed to get age from api: %v", err)
	}
	person.Age = age

	gender, err := r.api.GetPersonGenderByName(person.Gender)
	if err != nil {
		return domain.Person{}, fmt.Errorf("failed to get age from api: %v", err)
	}
	person.Gender = gender

	nationality, err := r.api.GetPersonNationalityByName(person.Nationality)
	if err != nil {
		return domain.Person{}, fmt.Errorf("failed to get nationality from api: %v", err)
	}
	person.Nationality = nationality

	id, err := r.repo.AddPerson(person)
	if err != nil {
		return domain.Person{}, fmt.Errorf("failed to add person: %v", err)
	}
	person.Id = id

	return person, nil
}

func (r *PersonService) DeletePersonById(id int) error {
	return r.repo.DeletePersonById(id)
}

func (r *PersonService) UpdatePersonById(person domain.Person) error {
	return r.repo.UpdatePersonById(person)
}
