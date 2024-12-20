package application

import (
	"context"
	er "git.ai-space.tech/coursework/backend/internal/domain/event/repository"
	"git.ai-space.tech/coursework/backend/internal/domain/sponsor/entity"
	sr "git.ai-space.tech/coursework/backend/internal/domain/sponsor/repository"
	"git.ai-space.tech/coursework/backend/internal/presentation/http/sponsors/sponsor_dto"
)

type SponsorService struct {
	SponsorRepo sr.SponsorRepository
	EventRepo   er.EventRepository
	ctx         context.Context
}

func NewSponsorService(sponsorRepo sr.SponsorRepository, eventRepo er.EventRepository) *SponsorService {
	return &SponsorService{
		SponsorRepo: sponsorRepo,
		EventRepo:   eventRepo,
	}
}

func (s *SponsorService) Create(dto *sponsor_dto.CreateSponsor) error {
	if _, err := s.GetSponsor(dto.Name); err == nil {
		return err
	}

	event, err := s.EventRepo.GetEventByName(dto.EventName)

	if err != nil {
		return err
	}

	sponsor := entity.NewSponsor(dto.Name, dto.ContactInfo, event.ID, dto.ContributionAmount)

	return s.SponsorRepo.CreateSponsor(sponsor)
}

func (s *SponsorService) GetSponsor(sponsorName string) (*entity.Sponsor, error) {
	return s.SponsorRepo.GetSponsorByName(sponsorName)
}

func (s *SponsorService) DeleteSponsor(sponsorName string) error {
	sponsor, err := s.GetSponsor(sponsorName)

	if err != nil {
		return err
	}

	return s.SponsorRepo.DeleteSponsor(sponsor.ID)
}

func (s *SponsorService) UpdateSponsor(sponsor *entity.Sponsor) error {

	return s.SponsorRepo.UpdateSponsor(sponsor)
}

func (s *SponsorService) GetAllSponsors() ([]entity.Sponsor, error) {

	return s.SponsorRepo.GetAllSponsors()
}
