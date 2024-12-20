package postgres

import (
	"context"
	"git.ai-space.tech/coursework/backend/internal/domain/sponsor/entity"
	"git.ai-space.tech/coursework/backend/internal/infrastructure"
	"github.com/jackc/pgx/v5"
)

type PgSponsorRepository struct {
	conn *pgx.Conn
	ctx  context.Context
}

func NewSponsorRepository(connections *infrastructure.Connections) *PgSponsorRepository {
	return &PgSponsorRepository{
		conn: connections.PostgresConnection,
		ctx:  connections.Ctx,
	}
}

func (s *PgSponsorRepository) CreateSponsor(sponsor *entity.Sponsor) error {
	query := `INSERT INTO sponsors (sponsor_name, contact_info, contribution_amount, event_id)
			  VALUES ($1, $2, $3, $4)`

	_, err := s.conn.Exec(s.ctx, query, sponsor.Name, sponsor.ContactInfo, sponsor.ContributionAmount, sponsor.EventID)

	if err != nil {
		return err
	}
	return nil
}

func (s *PgSponsorRepository) GetSponsorByName(sponsorName string) (*entity.Sponsor, error) {
	query := `SELECT sponsor_id, sponsor_name, contact_info, contribution_amount, events.event_id
			  FROM sponsors left join events using (event_id)
			  WHERE sponsor_name = $1`

	row := s.conn.QueryRow(s.ctx, query, sponsorName)

	var sponsor entity.Sponsor

	err := row.Scan(
		&sponsor.ID,
		&sponsor.Name,
		&sponsor.ContactInfo,
		&sponsor.ContributionAmount,
		&sponsor.EventID,
	)

	if err != nil {
		return nil, err
	}

	return &sponsor, nil
}

func (s *PgSponsorRepository) DeleteSponsor(id int) error {
	query := `DELETE FROM sponsors WHERE sponsor_id = $1`

	_, err := s.conn.Exec(s.ctx, query, id)

	if err != nil {
		return err
	}
	return nil
}

func (s *PgSponsorRepository) GetAllSponsors() ([]entity.Sponsor, error) {
	query := `SELECT sponsor_id, sponsor_name, contact_info, contribution_amount, events.event_id
			  FROM sponsors left join events using (event_id)`

	rows, err := s.conn.Query(s.ctx, query)

	if err != nil {
		return nil, err
	}

	var sponsors []entity.Sponsor

	for rows.Next() {
		var sponsor entity.Sponsor

		err = rows.Scan(
			&sponsor.ID,
			&sponsor.Name,
			&sponsor.ContactInfo,
			&sponsor.ContributionAmount,
			&sponsor.EventID,
		)

		if err != nil {
			return nil, err
		}

		sponsors = append(sponsors, sponsor)
	}

	return sponsors, nil
}

func (s *PgSponsorRepository) UpdateSponsor(sponsor *entity.Sponsor) error {
	query := `UPDATE sponsors SET sponsor_name = $1, contact_info = $2, contribution_amount = $3, event_id = $4
			  WHERE sponsor_id = $5`
	_, err := s.conn.Exec(s.ctx, query, sponsor.Name, sponsor.ContactInfo, sponsor.ContributionAmount, sponsor.EventID, sponsor.ID)

	if err != nil {
		return err
	}

	return nil
}
