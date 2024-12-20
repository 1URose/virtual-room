package repository

import (
	"git.ai-space.tech/coursework/backend/internal/domain/ticket/entity"
)

type TicketRepository interface {
	CreateTicket(ticket *entity.Ticket) error
	GetTicketByUserLoginAndEventName(userLogin string, eventName string) (*entity.Ticket, error)
	GetAllTickets() ([]entity.Ticket, error)
	DeleteTicket(userId int, eventId int) error
	UpdateTicket(ticket *entity.Ticket) error
}
