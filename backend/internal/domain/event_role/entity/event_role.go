package entity

type EventRole int

const (
	EventRoleParticipant EventRole = iota
	EventRoleSpeaker
	EventRoleModerator
)

func (er EventRole) ToString() string {
	switch er {
	case EventRoleParticipant:
		return "participant"
	case EventRoleSpeaker:
		return "speaker"
	case EventRoleModerator:
		return "moderator"
	default:
		return ""
	}
}
