package domain

import "errors"

var (
	AgeNotFound         = errors.New("Age not found")
	GenderNotFound      = errors.New("Gender not found")
	NationalityNotFound = errors.New("Nationality not found")
	NotFoundByParams    = errors.New("Not found by params")
	NotFoundById        = errors.New("Not found by id")
	NothingWasDeleted   = errors.New("Nothing was deleted")
	NothingWasUpdated   = errors.New("Nothing was updated")
)
