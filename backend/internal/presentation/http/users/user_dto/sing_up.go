package user_dto

import "git.ai-space.tech/coursework/backend/internal/domain/user_role/entity"

type SignUp struct {
	Name     string          `json:"name"`
	Email    string          `json:"login"`
	Password string          `json:"password"`
	Role     entity.UserRole `json:"user_role"`
}
