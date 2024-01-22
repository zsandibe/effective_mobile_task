package repository

import (
	"fmt"

	"github.com/zsandibe/effective_mobile_task/internal/domain"
)

func (r *PersonSql) FindPersonsBySearch(searchParams domain.SearchParams) ([]domain.Person, error) {
	var persons []domain.Person
	query := `
		SELECT * FROM persons WHERE Gender = ? AND Nationality = ?
		ORDER BY ID ASC
		OFFSET $3
		LIMIT $4
	`
	rows, err := r.db.Query(query, searchParams.PageStart, searchParams.PageLimit)
	if err != nil {
		return persons, err
	}
	defer rows.Close()
	for rows.Next() {
		var person domain.Person
		if err := rows.Scan(&person.Id, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Nationality); err != nil {
			return nil, fmt.Errorf("can`t find persons by name: %v", err)
		}
		persons = append(persons, person)
	}
	return persons, nil
}

func (r *PersonSql) FindPersonById(personId int) (domain.Person, error) {
	var person domain.Person
	query := `
		SELECT * FROM persons WHERE ID = ?
	`
	if err := r.db.QueryRow(query, personId).Scan(&person.Id, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Gender, &person.Nationality); err != nil {
		return domain.Person{}, fmt.Errorf("can`t find person by ID: %v", err)
	}

	return person, nil
}

func (r *PersonSql) AddPerson(person domain.Person) (int, error) {
	var id int
	query := `
		INSERT INTO persons (Name,Surname,Patronymic,Age,Gender,Nationality) VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING ID 
	`
	if err := r.db.QueryRow(query, &person.Id, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Gender, &person.Nationality).Scan(&id); err != nil {
		return 0, fmt.Errorf("can`t add person: %v", err)
	}
	return id, nil
}

func (r *PersonSql) DeletePersonById(personId int) error {
	query := `
		DELETE FROM persons WHERE ID = ?
	`
	if _, err := r.db.Exec(query, personId); err != nil {
		return fmt.Errorf("can`t delete person: %v", err)
	}
	return nil
}

func (r *PersonSql) UpdatePersonById(person domain.Person) error {
	query := `
	UPDATE persons SET Name = $1,
	Surname = $2,
	Patronymic = $3,
	Age = $4,
	Gender = $5,
	Nationality = $6
	WHERE ID = ?
	`

	if _, err := r.db.Exec(query, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Gender, &person.Nationality); err != nil {
		return fmt.Errorf("can`t update person: %v", err)
	}

	return nil
}
