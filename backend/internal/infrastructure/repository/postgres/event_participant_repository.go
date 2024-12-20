package postgres

import (
	"context"
	"errors"
	"fmt"
	"git.ai-space.tech/coursework/backend/internal/domain/event/entity"
	entity2 "git.ai-space.tech/coursework/backend/internal/domain/user/entity"
	"git.ai-space.tech/coursework/backend/internal/infrastructure"
	"github.com/jackc/pgx/v5"
)

type PgEventParticipantRepository struct {
	conn *pgx.Conn
	ctx  context.Context
}

func NewEventParticipantRouter(connections *infrastructure.Connections) *PgEventParticipantRepository {
	return &PgEventParticipantRepository{
		conn: connections.PostgresConnection,
		ctx:  connections.Ctx,
	}
}

func (r *PgEventParticipantRepository) BecomeParticipant(event *entity.Event, userId int) error {
	query := `
		INSERT INTO event_participants (event_id, user_id, role_in_event)
		VALUES ($1, $2, 'particpant')
		ON CONFLICT (event_id, user_id) DO NOTHING
	`
	_, err := r.conn.Exec(r.ctx, query, event.ID, userId)
	if err != nil {
		return fmt.Errorf("failed to add participant: %w", err)
	}
	return nil
}

func (r *PgEventParticipantRepository) GetAllEventParticipants(eventId int) ([]*entity2.User, error) {
	query := `
		SELECT u.user_id, u.username, u.email
		FROM users u
		INNER JOIN event_participants ep ON u.user_id = ep.user_id
		WHERE ep.event_id = $1
	`
	rows, err := r.conn.Query(r.ctx, query, eventId)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve participants: %w", err)
	}
	defer rows.Close()

	var users []*entity2.User
	for rows.Next() {
		var user entity2.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, &user)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", rows.Err())
	}

	return users, nil
}
func (r *PgEventParticipantRepository) GetAllEventsByParticpiantId(userId int) ([]*entity.Event, error) {

	query := `SELECT event_id, event_name, description, users.user_id, start_time, end_time, virtual_rooms.room_id
			  FROM events left join users using (user_id) left join virtual_rooms using (room_id)
			  WHERE users.user_id = $1`

	rows, err := r.conn.Query(r.ctx, query, userId)
	defer rows.Close()

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve events: %w", err)
	}

	var events []*entity.Event
	for rows.Next() {
		var event entity.Event
		var err = rows.Scan(&event.ID, &event.EventName, &event.Description, &event.OrganizerID, &event.StartTime, &event.EndTime, &event.VirtualRoomID)

		if err != nil {
			return nil, fmt.Errorf("failed to scan event: %w", err)
		}

		events = append(events, &event)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", rows.Err())
	}

	return events, nil
}

func (r *PgEventParticipantRepository) LeaveEvent(event *entity.Event, userId int) error {
	query := `
		DELETE FROM event_participants
		WHERE event_id = $1 AND user_id = $2
	`
	commandTag, err := r.conn.Exec(r.ctx, query, event.ID, userId)
	if err != nil {
		return fmt.Errorf("failed to remove participant: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("no participant found to remove")
	}

	return nil
}
