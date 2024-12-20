package postgres

import (
	"errors"
	"fmt"
	"git.ai-space.tech/coursework/backend/internal/domain/event/entity"
	"git.ai-space.tech/coursework/backend/internal/infrastructure"
	"git.ai-space.tech/coursework/backend/internal/infrastructure/repository/postgres/pg_errors"
	"github.com/jackc/pgx/v5"
	"golang.org/x/net/context"
)

type PgEventRepository struct {
	conn *pgx.Conn
	ctx  context.Context
}

func NewEventRepository(connections *infrastructure.Connections) *PgEventRepository {
	return &PgEventRepository{
		conn: connections.PostgresConnection,
		ctx:  connections.Ctx,
	}
}

func (r *PgEventRepository) CreateEvent(event *entity.Event) error {
	query := `INSERT INTO events (organizer_id, virtual_room_id, event_name, description, start_time, end_time)
			  VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.conn.Exec(r.ctx, query, event.OrganizerID, event.VirtualRoomID, event.EventName, event.Description, event.StartTime, event.EndTime)

	if err != nil {
		return fmt.Errorf("error creating event: %w", err)
	}

	return nil
}

func (r *PgEventRepository) GetEventByName(name string) (*entity.Event, error) {

	query := `SELECT event_id, event_name, description, users.user_id, start_time, end_time, virtual_rooms.room_id
			  FROM events left join users on events.organizer_id = users.user_id left join virtual_rooms on events.virtual_room_id = virtual_rooms.room_id
			  WHERE event_name = $1`

	row := r.conn.QueryRow(r.ctx, query, name)

	var event entity.Event

	err := row.Scan(
		&event.ID,
		&event.EventName,
		&event.Description,
		&event.OrganizerID,
		&event.StartTime,
		&event.EndTime,
		&event.VirtualRoomID,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pg_errors.NewEventNotFoundError(name)
		}

		return nil, err
	}

	return &event, nil

}

func (r *PgEventRepository) GetAllEvents() ([]entity.Event, error) {

	query := `SELECT event_id, event_name, description, events.organizer_id, start_time, end_time, events.virtual_room_id
			  FROM events
			      left join users on events.organizer_id=users.user_id
			      left join virtual_rooms vr on events.virtual_room_id=vr.room_id`

	rows, err := r.conn.Query(r.ctx, query)

	if err != nil {
		return nil, err
	}

	var events []entity.Event
	for rows.Next() {
		var event entity.Event
		err = rows.Scan(
			&event.ID,
			&event.EventName,
			&event.Description,
			&event.OrganizerID,
			&event.StartTime,
			&event.EndTime,
			&event.VirtualRoomID,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", rows.Err())
	}

	return events, nil
}

func (r *PgEventRepository) UpdateEvent(event *entity.Event) error {
	query := `UPDATE events SET event_name = $1, description = $2, start_time = $3, end_time = $4 WHERE event_id = $5`

	_, err := r.conn.Exec(r.ctx, query, event.EventName, event.Description, event.StartTime, event.EndTime, event.ID)

	if err != nil {
		return fmt.Errorf("error updating event: %w", err)
	}

	return nil
}

func (r *PgEventRepository) DeleteEvent(name string) error {
	query := `DELETE FROM events WHERE event_name = $1`

	_, err := r.conn.Exec(r.ctx, query, name)

	if err != nil {
		return fmt.Errorf("error deleting event: %w", err)
	}

	return nil
}
