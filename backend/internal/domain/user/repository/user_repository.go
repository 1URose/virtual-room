package repository

import "git.ai-space.tech/coursework/backend/internal/domain/user/entity"

type UserRepository interface {
	CreateUser(User *entity.User) error
	GetByEmail(login string) (*entity.User, error)
	GetAllUsers() ([]entity.User, error)
	UpdateUser(User *entity.User) error
	DeleteUser(login string) error
}
