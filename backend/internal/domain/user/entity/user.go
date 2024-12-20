package entity

import (
	"git.ai-space.tech/coursework/backend/internal/domain/user_role/entity"
	"time"
)

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string
	Role         entity.UserRole `json:"user_role"`
	DateCreated  time.Time       `json:"date_created"`
}

func NewUser(name, email, passwordHash string, role entity.UserRole) *User {
	return &User{
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
		Role:         role,
		DateCreated:  time.Now(),
	}
}

func (u *User) FromString(role string) entity.UserRole {
	switch role {
	case "admin":
		return entity.Admin
	case "organizer":
		return entity.Organizer
	case "participant":
		return entity.Participant
	default:
		return entity.UserRole(-1)
	}
}
