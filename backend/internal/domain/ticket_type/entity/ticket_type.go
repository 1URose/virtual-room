package entity

type TicketType int

const (
	TicketTypeStandard TicketType = iota
	TicketTypeVIP
)

func (tt TicketType) ToString() string {
	switch tt {
	case TicketTypeStandard:
		return "standard"
	case TicketTypeVIP:
		return "VIP"
	default:
		return ""
	}
}
