package entity

import "git.ai-space.tech/coursework/backend/internal/domain/event_role/entity"

type EventParticipant struct {
	EventID     int              `json:"event_id"`
	UserID      int              `json:"user_id"`
	RoleInEvent entity.EventRole `json:"role_in_event"`
}
