package jwt_manager

import (
	"errors"
	"time"

	"visualizer-go/internal/config"
	"visualizer-go/internal/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtManager struct {
	signingKey      []byte
	accessTokenTTL  uint
	refreshTokenTTL uint
}

type JwtClaims struct {
	UserID   uuid.UUID `json:"userId"`
	UserRole string    `json:"userRole"`
	jwt.RegisteredClaims
}

func NewJwtManager(cfg config.Jwt) *JwtManager {
	return &JwtManager{
		signingKey:      []byte(cfg.SigningKey),
		accessTokenTTL:  cfg.AccessTokenTTL,
		refreshTokenTTL: cfg.RefreshTokenTTL,
	}
}

func (m *JwtManager) GenerateTokens(user *models.User) (accessToken string, refreshToken string, err error) {
	accessToken, err = m.generateToken(user, m.accessTokenTTL)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = m.generateToken(user, m.refreshTokenTTL)
	if err != nil {
		return "", "", err
	}

	return
}

func (m *JwtManager) generateToken(user *models.User, ttl uint) (string, error) {
	claims := &JwtClaims{
		UserID:   user.ID,
		UserRole: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(ttl) * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(m.signingKey)
}

func (m *JwtManager) ParseToken(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return m.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func (m *JwtManager) ExtractUserID(tokenString string) (uuid.UUID, error) {
	claims, err := m.ParseToken(tokenString)
	if err != nil {
		return uuid.Nil, err
	}
	return claims.UserID, nil
}

func (m *JwtManager) ExtractUserRole(tokenString string) (string, error) {
	claims, err := m.ParseToken(tokenString)
	if err != nil {
		return "", err
	}
	return claims.UserRole, nil
}
