package application

import (
	"context"
	"fmt"
	"git.ai-space.tech/coursework/backend/internal/domain/event/entity"
	er "git.ai-space.tech/coursework/backend/internal/domain/event/repository"
	ur "git.ai-space.tech/coursework/backend/internal/domain/user/repository"
	vr "git.ai-space.tech/coursework/backend/internal/domain/virtual_room/repository"
	"git.ai-space.tech/coursework/backend/internal/presentation/http/events/event_dto"
)

type EventService struct {
	VirtualRoomRepository vr.VirtualRoomRepository
	UserRepository        ur.UserRepository
	EventRepository       er.EventRepository
	ctx                   context.Context
}

func NewEventService(eventRepo er.EventRepository, UserRepo ur.UserRepository, VirtualRepo vr.VirtualRoomRepository) *EventService {
	return &EventService{
		EventRepository:       eventRepo,
		UserRepository:        UserRepo,
		VirtualRoomRepository: VirtualRepo,
	}
}

func (es *EventService) Create(dto event_dto.CreateEvent) error {
	if _, err := es.EventRepository.GetEventByName(dto.Title); err == nil {
		return fmt.Errorf("event with name '%s' already exists", dto.Title)
	}

	virtualRoom, err := es.VirtualRoomRepository.GetByName(dto.VirtualRoomName)

	if err != nil {
		return err
	}

	organizer, err := es.UserRepository.GetByEmail(dto.OrganizerLogin)

	if err != nil {
		return err
	}

	event := entity.NewEvent(dto.Title, dto.Description, dto.StartDate, dto.EndDate, organizer.ID, virtualRoom.ID)

	return es.EventRepository.CreateEvent(event)
}

func (es *EventService) GetEventByName(name string) (*entity.Event, error) {
	return es.EventRepository.GetEventByName(name)
}

func (es *EventService) GetAllEvents() ([]entity.Event, error) {
	return es.EventRepository.GetAllEvents()
}

func (es *EventService) DeleteEvent(name string) error {

	return es.EventRepository.DeleteEvent(name)
}

func (es *EventService) UpdateEvent(event *entity.Event) error {

	return es.EventRepository.UpdateEvent(event)
}
