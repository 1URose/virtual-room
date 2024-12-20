package postgres

import (
	"context"
	"fmt"
	"git.ai-space.tech/coursework/backend/internal/domain/ticket/entity"
	"git.ai-space.tech/coursework/backend/internal/infrastructure"
	"github.com/jackc/pgx/v5"
)

type PgTicketRepository struct {
	conn *pgx.Conn
	ctx  context.Context
}

func NewTicketRepository(connections *infrastructure.Connections) *PgTicketRepository {
	return &PgTicketRepository{
		conn: connections.PostgresConnection,
		ctx:  connections.Ctx,
	}
}

func (tr *PgTicketRepository) CreateTicket(ticket *entity.Ticket) error {
	query := `INSERT INTO tickets (event_id, user_id, ticket_type, price, purchase_date)
			  VALUES ($1, $2, $3, $4, $5)`

	_, err := tr.conn.Exec(tr.ctx, query, ticket.EventID, ticket.UserID, ticket.TicketType.ToString(), ticket.Price, ticket.PurchaseDate)

	if err != nil {
		return err
	}
	return nil
}

func (tr *PgTicketRepository) GetTicketByUserLoginAndEventName(userLogin string, eventName string) (*entity.Ticket, error) {
	query := `SELECT ticket_id, events.event_id, users.user_id, ticket_type, price, purchase_date
			  FROM tickets left join events on tickets.event_id = events.event_id  left join users on tickets.user_id = users.user_id
			  WHERE users.email = $1 AND event_name = $2`

	fmt.Printf("query: %s, userLogin: %s, eventName: %s\n", query, userLogin, eventName)
	row := tr.conn.QueryRow(tr.ctx, query, userLogin, eventName)

	var ticketType string

	var ticket entity.Ticket

	err := row.Scan(
		&ticket.ID,
		&ticket.EventID,
		&ticket.UserID,
		&ticketType,
		&ticket.Price,
		&ticket.PurchaseDate,
	)

	ticket.TicketType = entity.TicketTypeFromString(ticketType)

	if err != nil {
		return nil, err
	}

	return &ticket, nil

}

func (tr *PgTicketRepository) GetAllTickets() ([]entity.Ticket, error) {
	query := `SELECT ticket_id, events.event_id, users.user_id, ticket_type, price, purchase_date
			  FROM tickets left join events using (event_id) left join users using (user_id)`

	rows, err := tr.conn.Query(tr.ctx, query)

	if err != nil {
		return nil, err
	}

	var ticketType string

	var tickets []entity.Ticket

	for rows.Next() {
		var ticket entity.Ticket
		err := rows.Scan(
			&ticket.ID,
			&ticket.EventID,
			&ticket.UserID,
			&ticketType,
			&ticket.Price,
			&ticket.PurchaseDate,
		)

		ticket.TicketType = entity.TicketTypeFromString(ticketType)
		if err != nil {
			return nil, err
		}

		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

func (tr *PgTicketRepository) DeleteTicket(userId int, eventId int) error {
	query := `DELETE FROM tickets WHERE tickets.user_id = $1 AND tickets.event_id = $2`

	_, err := tr.conn.Exec(tr.ctx, query, userId, eventId)

	if err != nil {
		return err
	}
	return nil
}

func (tr *PgTicketRepository) UpdateTicket(ticket *entity.Ticket) error {
	query := `UPDATE tickets
			  SET event_id = $1, user_id = $2, ticket_type = $3, price = $4, purchase_date = $5
			  WHERE ticket_id = $6`

	_, err := tr.conn.Exec(tr.ctx, query, ticket.EventID, ticket.UserID, ticket.TicketType, ticket.Price, ticket.PurchaseDate, ticket.ID)

	if err != nil {
		return err
	}

	return nil
}
