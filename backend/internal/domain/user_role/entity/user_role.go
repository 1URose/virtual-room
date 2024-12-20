package entity

type UserRole int

const (
	Admin UserRole = iota
	Organizer
	Participant
)

func (r UserRole) ToString() string {
	switch r {
	case Admin:
		return "admin"
	case Organizer:
		return "organizer"
	case Participant:
		return "participant"
	default:
		return ""
	}
}

func RoleFromString(role string) UserRole {
	switch role {
	case "admin":
		return 0
	case "organizer":
		return 1
	case "participant":
		return 2
	default:
		return -1
	}
}
