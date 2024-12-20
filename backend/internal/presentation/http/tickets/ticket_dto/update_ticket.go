package ticket_dto

import "git.ai-space.tech/coursework/backend/internal/domain/ticket_type/entity"

type UpdateTicket struct {
	EventName  string            `json:"event_name"`
	UserLogin  string            `json:"login"`
	TicketType entity.TicketType `json:"ticket_type"`
	Price      float64           `json:"price"`
}
