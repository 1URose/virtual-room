package repository

import (
	"git.ai-space.tech/coursework/backend/internal/domain/event/entity"
)

type EventRepository interface {
	CreateEvent(event *entity.Event) error
	GetEventByName(name string) (*entity.Event, error)
	GetAllEvents() ([]entity.Event, error)
	UpdateEvent(event *entity.Event) error
	DeleteEvent(name string) error
}
