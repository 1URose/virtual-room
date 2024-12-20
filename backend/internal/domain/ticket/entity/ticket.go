package entity

import (
	"git.ai-space.tech/coursework/backend/internal/domain/ticket_type/entity"
	"time"
)

type Ticket struct {
	ID           int               `json:"id"`
	EventID      int               `json:"event_id"`
	UserID       int               `json:"user_id"`
	TicketType   entity.TicketType `json:"ticket_type"`
	Price        float64           `json:"price"`
	PurchaseDate time.Time         `json:"purchase_date"`
}

func NewTicket(eventID int, userID int, ticketType entity.TicketType, price float64, purchaseDate time.Time) *Ticket {
	return &Ticket{
		EventID:      eventID,
		UserID:       userID,
		TicketType:   ticketType,
		Price:        price,
		PurchaseDate: purchaseDate,
	}
}

func TicketTypeFromString(ticketType string) entity.TicketType {
	switch ticketType {
	case "standard":
		return entity.TicketTypeStandard
	case "VIP":
		return entity.TicketTypeVIP
	default:
		return entity.TicketType(-1)
	}
}
