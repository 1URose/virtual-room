package application

import (
	"context"
	"fmt"
	"git.ai-space.tech/coursework/backend/internal/domain/event/entity"
	"git.ai-space.tech/coursework/backend/internal/domain/event_participant/repository"
	entity2 "git.ai-space.tech/coursework/backend/internal/domain/user/entity"
)

type EventParticipiantService struct {
	repo repository.EventParticipiantRepository
	ctx  context.Context
}

func NewEventParticipantService(repo repository.EventParticipiantRepository) *EventParticipiantService {
	return &EventParticipiantService{
		repo: repo,
	}
}

func (s *EventParticipiantService) BecomeParticipiant(event *entity.Event, userId int) error {
	err := s.repo.BecomeParticipant(event, userId)

	if err != nil {
		return fmt.Errorf("Error while becoming participiant: %+v\n", err)
	}

	return nil
}

func (s *EventParticipiantService) LeaveEvent(event *entity.Event, userId int) error {
	err := s.repo.LeaveEvent(event, userId)

	if err != nil {
		return fmt.Errorf("Error while leaving event: %+v\n", err)
	}

	return nil
}

func (s *EventParticipiantService) GetAllEventParticipants(event *entity.Event) ([]*entity2.User, error) {
	users, err := s.repo.GetAllEventParticipants(event.ID)
	if err != nil {
		return nil, fmt.Errorf("Error while getting all event participiants: %+v\n", err)
	}

	return users, nil
}

func (s *EventParticipiantService) GetAllEventsByParticipantId(userId int) ([]*entity.Event, error) {
	events, err := s.repo.GetAllEventsByParticpiantId(userId)
	if err != nil {
		return nil, fmt.Errorf("Error while getting all events by participiant: %+v\n", err)
	}

	return events, nil
}
