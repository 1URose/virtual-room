package postgres

import (
	"context"
	"errors"
	"git.ai-space.tech/coursework/backend/internal/domain/virtual_room/entity"
	"git.ai-space.tech/coursework/backend/internal/infrastructure"
	"git.ai-space.tech/coursework/backend/internal/infrastructure/repository/postgres/pg_errors"
	"github.com/jackc/pgx/v5"
)

type PgVirtualRoomRepository struct {
	Con *pgx.Conn
	Ctx context.Context
}

func NewVirtualRoomRepository(connections *infrastructure.Connections) *PgVirtualRoomRepository {
	return &PgVirtualRoomRepository{
		Con: connections.PostgresConnection,
		Ctx: connections.Ctx,
	}
}

func (r *PgVirtualRoomRepository) CreateVirtualRoom(room *entity.VirtualRoom) error {

	query := `INSERT INTO virtual_rooms (room_name, capacity) VALUES ($1, $2)`

	_, err := r.Con.Exec(r.Ctx, query, room.RoomName, room.Capacity)

	return err
}

func (r *PgVirtualRoomRepository) GetByName(roomName string) (*entity.VirtualRoom, error) {
	quere := `SELECT room_id, room_name, capacity FROM virtual_rooms WHERE room_name = $1`

	var room entity.VirtualRoom
	row := r.Con.QueryRow(r.Ctx, quere, roomName)

	err := row.Scan(&room.ID, &room.RoomName, &room.Capacity)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pg_errors.NewVirtualRoomNotFoundError(roomName)
		}

		return nil, err
	}

	return &room, nil
}

func (r *PgVirtualRoomRepository) GetAll() ([]entity.VirtualRoom, error) {
	query := `SELECT room_id, room_name, capacity FROM virtual_rooms`

	rows, err := r.Con.Query(r.Ctx, query)

	if err != nil {
		return nil, err
	}

	var rooms []entity.VirtualRoom
	for rows.Next() {
		var room entity.VirtualRoom
		err = rows.Scan(&room.ID, &room.RoomName, &room.Capacity)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil

}

func (r *PgVirtualRoomRepository) DeleteByName(name string) error {
	query := `DELETE FROM virtual_rooms WHERE room_name = $1`

	_, err := r.Con.Exec(r.Ctx, query, name)

	return err
}

func (r *PgVirtualRoomRepository) UpdateVirtualRoom(room *entity.VirtualRoom) error {
	query := `UPDATE virtual_rooms
			  SET room_name = coalesce($1, room_name), capacity = coalesce($2, capacity)
			  WHERE room_id = $3`

	_, err := r.Con.Exec(r.Ctx, query, room.RoomName, room.Capacity, room.ID)

	return err
}
