package personApi

import "github.com/zsandibe/effective_mobile_task/config"

type PersonApi interface {
	GetPersonAgeByName(name string) (int, error)
	GetPersonGenderByName(name string) (string, error)
	GetPersonNationalityByName(name string) (string, error)
}

type PersonApiService struct {
	config *config.Config
}

func NewPersonApiService(config *config.Config) *PersonApiService {
	return &PersonApiService{
		config: config,
	}
}
