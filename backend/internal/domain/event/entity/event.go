package entity

import "time"

type Event struct {
	ID            int       `json:"id"`
	EventName     string    `json:"event_name"`
	Description   string    `json:"description"`
	OrganizerID   int       `json:"organizer_id"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	VirtualRoomID int       `json:"virtual_room_id"`
}

func NewEvent(eventName string, description string, startTime time.Time, endTime time.Time, organizerID int, virtualRoomID int) *Event {
	return &Event{
		EventName:     eventName,
		Description:   description,
		OrganizerID:   organizerID,
		StartTime:     startTime,
		EndTime:       endTime,
		VirtualRoomID: virtualRoomID,
	}
}
