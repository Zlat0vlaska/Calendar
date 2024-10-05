package handler

import (
	"github.com/Zlat0vlaska/Calendar/internal/service"
)

type EventHandler struct {
	eventService service.EventService
}

func NewEventHandler(eventService service.EventService) *EventHandler {
	return &EventHandler{eventService: eventService}
}
