package application

import (
	"context"
	"fmt"
	er "git.ai-space.tech/coursework/backend/internal/domain/event/repository"
	"git.ai-space.tech/coursework/backend/internal/domain/ticket/entity"
	tr "git.ai-space.tech/coursework/backend/internal/domain/ticket/repository"

	ur "git.ai-space.tech/coursework/backend/internal/domain/user/repository"
	"git.ai-space.tech/coursework/backend/internal/presentation/http/tickets/ticket_dto"
	"time"
)

type TicketService struct {
	TicketRepo tr.TicketRepository
	EventRepo  er.EventRepository
	UserRepo   ur.UserRepository

	ctx context.Context
}

func NewTicketService(ticketRepo tr.TicketRepository, eventRepo er.EventRepository, userRepo ur.UserRepository) *TicketService {
	return &TicketService{TicketRepo: ticketRepo,
		EventRepo: eventRepo,
		UserRepo:  userRepo,
	}
}

func (ts *TicketService) CreateTicket(dto ticket_dto.CreateTicket) error {

	if _, err := ts.GetTicket(dto.UserLogin, dto.EventName); err == nil {

		return fmt.Errorf("билет на событие '%s' человеком '%s' уже куплен", dto.UserLogin, dto.EventName)
	}

	event, err := ts.EventRepo.GetEventByName(dto.EventName)

	if err != nil {

		return fmt.Errorf("событие '%s' не найдено", dto.EventName)
	}

	user, err := ts.UserRepo.GetByEmail(dto.UserLogin)

	if err != nil {

		return fmt.Errorf("пользователь '%s' не найден", dto.UserLogin)
	}

	ticket := entity.NewTicket(event.ID, user.ID, dto.TicketType, dto.Price, time.Now())

	return ts.TicketRepo.CreateTicket(ticket)
}

func (ts *TicketService) GetTicket(userLogin, eventName string) (*entity.Ticket, error) {

	return ts.TicketRepo.GetTicketByUserLoginAndEventName(userLogin, eventName)
}

func (ts *TicketService) GetAllTickets() ([]entity.Ticket, error) {

	return ts.TicketRepo.GetAllTickets()
}

func (ts *TicketService) DeleteTicket(userLogin string, eventName string) error {

	if _, err := ts.GetTicket(userLogin, eventName); err != nil {

		return err
	}

	user, err := ts.UserRepo.GetByEmail(userLogin)

	if err != nil {

		return fmt.Errorf("пользователь '%s' не найден", userLogin)

	}

	event, err := ts.EventRepo.GetEventByName(eventName)

	if err != nil {

		return fmt.Errorf("событие '%s' не найдено", eventName)
	}

	return ts.TicketRepo.DeleteTicket(user.ID, event.ID)
}

func (ts *TicketService) UpdateTicket(ticket *entity.Ticket) error {

	return ts.TicketRepo.UpdateTicket(ticket)
}
