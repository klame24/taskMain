package services

import (
	"context"
	"errors"
	"taskMain/internal/auth/jwt"
	"taskMain/internal/auth/password"
	authdto "taskMain/internal/dto/authDTO"
	"taskMain/internal/dto/userDTO"
	"taskMain/internal/models"
	"taskMain/internal/repositories"
	"time"
)

type AuthService interface {
	Register(ctx context.Context, req userDTO.CreateUserRequest) (int, error)
	Login(ctx context.Context, req authdto.LoginRequest) (*authdto.LoginResponse, error)
}

type authService struct {
	r               repositories.UserRepository
	tokenRepo       *jwt.TokenRepository
	jwtManager      *jwt.Manager
	refreshTokenExp time.Duration
}

func NewAuthService(
	r repositories.UserRepository, tokenRepo *jwt.TokenRepository,
	jwtManager *jwt.Manager, refreshTokenExtp time.Duration,
) AuthService {
	return &authService{
		r:               r,
		tokenRepo:       tokenRepo,
		jwtManager:      jwtManager,
		refreshTokenExp: refreshTokenExtp,
	}
}

func (s *authService) Register(ctx context.Context, req userDTO.CreateUserRequest) (int, error) {
	hashedPassword, err := password.HashPassword(req.Password)
	if err != nil {
		return 0, err
	}

	user := &models.User{
		Name:         req.Name,
		Surname:      req.Surname,
		Nickname:     req.Nickname,
		Email:        req.Email,
		PasswordHash: hashedPassword,
	}

	return s.r.Create(ctx, user)
}

func (s *authService) Login(ctx context.Context, req authdto.LoginRequest) (*authdto.LoginResponse, error) {
	user, err := s.r.GetByNickname(ctx, req.Nickname)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if !password.VerifyPassword(req.Password, user.PasswordHash) {
		return nil, errors.New("invalid password")
	}

	accessToken, err := s.jwtManager.GenerateAccessToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	refreshTokenExpiresAt := time.Now().Add(s.refreshTokenExp)
	err = s.tokenRepo.SaveRefreshToken(ctx, user.ID, refreshToken, refreshTokenExpiresAt)
	if err != nil {
		return nil, err
	}

	return &authdto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int(s.jwtManager.AccessTokenExp.Seconds()),
		TokenType:    "Bearer",
		User: userDTO.GetUserResponse{
			Name:     user.Name,
			Surname:  user.Surname,
			Nickname: user.Nickname,
			Email:    user.Email,
		},
	}, nil
}
