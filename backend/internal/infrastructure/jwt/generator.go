package jwt

import (
	"git.ai-space.tech/coursework/backend/internal/domain/user_role/entity"
	"github.com/golang-jwt/jwt"
	"time"
)

type Claims struct {
	Role string `json:"roles"`
	jwt.StandardClaims
}

const AccessTTL = 60 * 30 * time.Second
const RefreshTTL = AccessTTL * 2
const jwtSecret = "mdsadmkmksmk!dsmadsmaim@#!#!23123dx"

func GenerateAccessToken(email string, userRole entity.UserRole) (string, error) {
	claims := &Claims{
		Role: userRole.ToString(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func GenerateRefreshToken(email string, userRole entity.UserRole) (string, error) {
	claims := &Claims{
		Role: userRole.ToString(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(RefreshTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}
