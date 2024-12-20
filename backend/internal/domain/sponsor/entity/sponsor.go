package entity

type Sponsor struct {
	ID                 int     `json:"id"`
	Name               string  `json:"name"`
	ContactInfo        string  `json:"contact_info"`
	EventID            int     `json:"event_id"`
	ContributionAmount float64 `json:"contribution_amount"`
}

func NewSponsor(name string, contactInfo string, eventID int, contributionAmount float64) *Sponsor {
	return &Sponsor{
		Name:               name,
		ContactInfo:        contactInfo,
		EventID:            eventID,
		ContributionAmount: contributionAmount,
	}
}
