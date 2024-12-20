package pg_errors

type TicketNotFoundError struct {
	userLogin string
	eventName string
}

func NewTicketNotFoundError(userLogin string, eventName string) *TicketNotFoundError {
	return &TicketNotFoundError{userLogin: userLogin, eventName: eventName}
}

func (e *TicketNotFoundError) Error() string {
	return "Ticket not found " + e.userLogin + "enevt = " + e.eventName
}
