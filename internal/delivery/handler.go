package delivery

import (
	personService "github.com/zsandibe/effective_mobile_task/internal/service"
)

type Handler struct {
	service *personService.Service
}

func NewHandler(service *personService.Service) *Handler {
	return &Handler{service: service}
}
