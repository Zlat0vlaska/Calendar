package service

import (
	"project/internal/domain"

	"github.com/Zlat0vlaska/Calendar/internal/repository"
)

type EventService interface {
	CreateEvent(event *domain.Event) error
	UpdateEvent(event *domain.Event) error
	DeleteEvent(id int64) error
	GetEventByID(id int64) (*domain.Event, error)
	GetEventByUID(uid string) (*domain.Event, error)
}

type EventServiceImpl struct {
	repo repository.EventRepository
}

func NewEventService(repo repository.EventRepository) *EventServiceImpl {
	return &EventServiceImpl{repo: repo}
}
