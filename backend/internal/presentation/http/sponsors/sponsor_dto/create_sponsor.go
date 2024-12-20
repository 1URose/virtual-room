package sponsor_dto

type CreateSponsor struct {
	Name               string  `json:"name"`
	ContactInfo        string  `json:"contact_info"`
	EventName          string  `json:"event_name"`
	ContributionAmount float64 `json:"contribution_amount"`
}
