package event_participiant_dto

import "git.ai-space.tech/coursework/backend/internal/domain/event/entity"

type BecomeParticipantDto struct {
	Event  entity.Event `json:"event"`
	UserId int          `json:"user_id"`
}
