package person

import "github.com/zsandibe/effective_mobile_task/internal/domain"

func (r *PersonService) FindPersonsList(searchParams domain.SearchParams) ([]domain.Person, error) {
	return nil, nil
}

func (r *PersonService) FindPersonById(personId int) (domain.Person, error) {
	return domain.Person{}, nil
}

func (r *PersonService) AddPerson(person domain.Person) (int, error) {
	return 0, nil
}

func (r *PersonService) DeletePersonById(personId int) error {
	return nil
}

func (r *PersonService) UpdatePersonById(person domain.Person) error {
	return nil
}
