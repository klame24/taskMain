package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Manager struct {
	SecretKey      string
	AccessTokenExp time.Duration
}

func NewManager(secretkey string, accessTokenExp time.Duration) *Manager {
	return &Manager{
		SecretKey:      secretkey,
		AccessTokenExp: accessTokenExp,
	}
}

type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func (m *Manager) GenerateAccessToken(userID int, email string) (string, error) {
	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.AccessTokenExp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(m.SecretKey))
}

func (m *Manager) VerifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(m.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
