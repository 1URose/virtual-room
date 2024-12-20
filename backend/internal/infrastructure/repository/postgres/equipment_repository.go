package postgres

import (
	"context"
	"git.ai-space.tech/coursework/backend/internal/domain/equipment/entity"
	eu "git.ai-space.tech/coursework/backend/internal/domain/user/entity"
	"git.ai-space.tech/coursework/backend/internal/infrastructure"
	"github.com/jackc/pgx/v5"
)

type PgEquipmentRepository struct {
	conn *pgx.Conn
	ctx  context.Context
}

func NewEquipmentRepository(connections *infrastructure.Connections) *PgEquipmentRepository {
	return &PgEquipmentRepository{
		conn: connections.PostgresConnection,
		ctx:  connections.Ctx,
	}
}

func (er *PgEquipmentRepository) CreateEquipment(equipment *entity.Equipment) error {
	query := `INSERT INTO equipments (equipment_name, user_id) VALUES ($1, $2)`

	_, err := er.conn.Exec(er.ctx, query, equipment.Name, equipment.UserID)

	if err != nil {
		return err
	}
	return nil

}

func (er *PgEquipmentRepository) GetUserByEquipmentName(equipmentName string) (*eu.User, error) {
	query := `SELECT user_id, username, email, password_hash, role 
			  FROM equipments left join users using (user_id)
			  WHERE equipment_name = $1`

	row := er.conn.QueryRow(er.ctx, query, equipmentName)

	var user eu.User

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
