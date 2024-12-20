package repository

import (
	"git.ai-space.tech/coursework/backend/internal/domain/sponsor/entity"
)

type SponsorRepository interface {
	CreateSponsor(sponsor *entity.Sponsor) error
	GetSponsorByName(sponsorName string) (*entity.Sponsor, error)
	DeleteSponsor(id int) error
	UpdateSponsor(sponsor *entity.Sponsor) error
	GetAllSponsors() ([]entity.Sponsor, error)
}
