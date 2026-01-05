package services

import (
	"context"
	"taskMain/internal/models"
	"taskMain/internal/repositories"
)

type UserService interface {
	Create(
		ctx context.Context, name, surname,
		nickname, email, passwordHash string,
	) (int, error)
	GetByID(ctx context.Context, userID int) (*models.User, error)
}

type userService struct {
	r repositories.UserRepository
}

func NewUserService(r repositories.UserRepository) UserService {
	return &userService{
		r: r,
	}
}

func (s *userService) Create(
	ctx context.Context, name, surname,
	nickname, email, passwordHash string,
) (int, error) {
	user := models.User{
		Name:         name,
		Surname:      surname,
		Nickname:     nickname,
		Email:        email,
		PasswordHash: passwordHash,
	}

	userID, err := s.r.Create(ctx, &user)

	return userID, err
}

func (s *userService) GetByID(ctx context.Context, userID int) (*models.User, error) {
	user, err := s.r.GetByID(ctx, userID)

	return user, err
}
