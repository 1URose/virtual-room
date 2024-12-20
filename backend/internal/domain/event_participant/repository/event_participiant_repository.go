package repository

import (
	"git.ai-space.tech/coursework/backend/internal/domain/event/entity"
	entity2 "git.ai-space.tech/coursework/backend/internal/domain/user/entity"
)

type EventParticipiantRepository interface {
	BecomeParticipant(event *entity.Event, userId int) error
	GetAllEventParticipants(eventId int) ([]*entity2.User, error)
	GetAllEventsByParticpiantId(userId int) ([]*entity.Event, error)
	LeaveEvent(event *entity.Event, userId int) error
}
