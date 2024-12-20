package validator

import "git.ai-space.tech/coursework/backend/internal/infrastructure/http/ginConfig/routes/event_dto"

type TicketValidator struct {
	eDto event_dto.CreateEvent
}

func NewTicketValidator(eDto event_dto.CreateEvent) *TicketValidator {
	return &TicketValidator{eDto: eDto}
}
