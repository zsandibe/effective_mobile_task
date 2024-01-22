package domain

import "errors"

var (
	AgeNotFound         = errors.New("Age not found")
	GenderNotFound      = errors.New("Gender not found")
	NationalityNotFound = errors.New("Nationality not found")
)
