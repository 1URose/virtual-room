package event_dto

import "time"

type CreateEvent struct {
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	VirtualRoomName string    `json:"virtual_room_name"`
	OrganizerLogin  string    `json:"organizer_login"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
}
