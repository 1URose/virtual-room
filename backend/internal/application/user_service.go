package application

import (
	"context"
	"errors"
	"fmt"
	"git.ai-space.tech/coursework/backend/internal/domain/user/entity"
	"git.ai-space.tech/coursework/backend/internal/domain/user/repository"
	"git.ai-space.tech/coursework/backend/internal/infrastructure/cache/redis"
	"git.ai-space.tech/coursework/backend/internal/infrastructure/jwt"
	"git.ai-space.tech/coursework/backend/internal/infrastructure/jwt/jwt_errors"
	"git.ai-space.tech/coursework/backend/internal/infrastructure/repository/postgres/pg_errors"
	"git.ai-space.tech/coursework/backend/internal/presentation/http/users/user_dto"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo      repository.UserRepository
	redisConn *redis.Connection
	ctx       context.Context
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(user *entity.User) error {

	if _, err := s.repo.GetByEmail(user.Email); err == nil {
		return fmt.Errorf("пользователь с таким email уже существует")
	}

	return s.repo.CreateUser(user)
}

func (s *UserService) Auth(dto user_dto.Auth) (*user_dto.Tokens, error) {
	user, err := s.repo.GetByEmail(dto.Email)

	if err != nil {
		if errors.Is(err, pg_errors.NewUserNotFoundError(dto.Email)) {
			return nil, fmt.Errorf(err.Error())
		}

		return nil, fmt.Errorf("internal server error: %+v\n", err)
	}

	tokens, verifyError := s.verifyUserCreds(dto, user)

	if verifyError != nil {
		return nil, err
	}

	if tokens != nil {
		return tokens, nil
	}

	return nil, errors.New("такой пользователь не найден")
}

func (s *UserService) verifyUserCreds(dto user_dto.Auth, user *entity.User) (*user_dto.Tokens, error) {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(dto.Password))
	if err != nil {
		return nil, err
	}

	accessToken, err := jwt.GenerateAccessToken(user.Email, user.Role)

	if err != nil {
		if errors.Is(err, &jwt_errors.GenerateTokenError{}) {
			return nil, fmt.Errorf(err.Error())
		}
	}

	refreshToken, err := jwt.GenerateRefreshToken(user.Email, user.Role)

	if errors.Is(err, &jwt_errors.GenerateTokenError{}) {
		return nil, fmt.Errorf(err.Error())
	}

	return &user_dto.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}

func (s *UserService) SingUp(dto user_dto.SignUp) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user := entity.NewUser(dto.Name, dto.Email, string(hashedPassword), dto.Role)

	err = s.Create(user)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetAllUsers() ([]entity.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) GetUserByLogin(email string) (*entity.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *UserService) DeleteUser(email string) error {
	if _, err := s.repo.GetByEmail(email); err != nil {
		return fmt.Errorf("пользователь с таким email не существует")
	}

	return s.repo.DeleteUser(email)
}

func (s *UserService) UpdateUser(user *entity.User) error {

	if _, err := s.repo.GetByEmail(user.Email); err != nil {
		return fmt.Errorf("пользователь с таким email не существует")
	}

	return s.repo.UpdateUser(user)

}
