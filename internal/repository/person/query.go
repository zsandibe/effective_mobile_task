package repository

import "github.com/zsandibe/effective_mobile_task/internal/domain"

func (r *PersonSql) FindPersonsList(searchParams domain.SearchParams) ([]domain.Person, error) {
	return nil, nil
}

func (r *PersonSql) FindPersonById(personId int) (domain.Person, error) {
	return domain.Person{}, nil
}

func (r *PersonSql) AddPerson(person domain.Person) (int, error) {
	return 0, nil
}

func (r *PersonSql) DeletePersonById(personId int) error {
	return nil
}

func (r *PersonSql) UpdatePersonById(person domain.Person) error {
	return nil
}
