package pg_errors

type EventNotFoundError struct {
	Name string
}

func NewEventNotFoundError(name string) *EventNotFoundError {
	return &EventNotFoundError{Name: name}
}

func (e *EventNotFoundError) Error() string {
	return "Event not found " + e.Name
}
