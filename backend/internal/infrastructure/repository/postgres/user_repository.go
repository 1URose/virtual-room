package postgres

import (
	"context"
	"errors"
	"fmt"
	"git.ai-space.tech/coursework/backend/internal/domain/user/entity"
	"git.ai-space.tech/coursework/backend/internal/infrastructure"
	"git.ai-space.tech/coursework/backend/internal/infrastructure/repository/postgres/pg_errors"
	"github.com/jackc/pgx/v5"
)

type PgUserRepository struct {
	Con *pgx.Conn
	Ctx context.Context
}

func NewUserRepository(connections *infrastructure.Connections) *PgUserRepository {
	return &PgUserRepository{
		Con: connections.PostgresConnection,
		Ctx: connections.Ctx,
	}
}

func (r *PgUserRepository) CreateUser(user *entity.User) error {
	query := `INSERT INTO users (username, email, password_hash, role, date_created) VALUES ($1, $2, $3, $4, $5)`

	_, err := r.Con.Exec(r.Ctx, query, user.Name, user.Email, user.PasswordHash, user.Role.ToString(), user.DateCreated)

	if err != nil {
		return fmt.Errorf("error creating event: %w", err)
	}

	return err
}

func (r *PgUserRepository) GetByEmail(email string) (*entity.User, error) {
	fmt.Println(email)
	query := `
		SELECT user_id, username, email, password_hash, role, date_created
		FROM users
		WHERE email = $1
	`
	u := entity.User{}

	var user entity.User
	var role string

	row := r.Con.QueryRow(r.Ctx, query, email)
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&role,
		&user.DateCreated,
	)

	user.Role = u.FromString(role)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pg_errors.NewUserNotFoundError(email)
		}

		return nil, err
	}

	return &user, nil
}

func (r *PgUserRepository) GetAllUsers() ([]entity.User, error) {
	query := `
		SELECT user_id, username, email, password_hash, role, date_created
		FROM users
	`

	rows, err := r.Con.Query(r.Ctx, query)
	if err != nil {
		return nil, err
	}

	u := entity.User{}

	var users []entity.User

	var role string

	for rows.Next() {
		var user entity.User
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.PasswordHash,
			&role,
			&user.DateCreated,
		)
		user.Role = u.FromString(role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *PgUserRepository) UpdateUser(user *entity.User) error {
	query := `UPDATE users
			  SET username = $1, email = $2, password_hash = $3, role = $4, date_created = $5
			  WHERE user_id = $6`

	_, err := r.Con.Exec(r.Ctx, query, user.Name, user.Email, user.PasswordHash, user.Role.ToString(), user.DateCreated, user.ID)

	return err

}
func (r *PgUserRepository) DeleteUser(login string) error {

	query := `DELETE FROM users WHERE email = $1`

	_, err := r.Con.Exec(r.Ctx, query, login)

	return err
}
