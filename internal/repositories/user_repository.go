package repositories

import (
	"context"
	"taskMain/internal/models"

	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (int, error)
	GetByID(ctx context.Context, userID int) (*models.User, error)
}

type userRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) (int, error) {
	sqlQuery := `
		INSERT INTO users(name, surname, nickname, email, password_hash)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	var userID int

	err := r.db.QueryRow(ctx, sqlQuery,
		user.Name,
		user.Surname,
		user.Nickname,
		user.Email,
		user.PasswordHash).Scan(&userID)

	return userID, err
}

func (r *userRepository) GetByID(ctx context.Context, userID int) (*models.User, error) {
	user := models.User{}

	sqlQuery := `
		SELECT 
			name, surname, nickname, email
		FROM
			users
		WHERE users.id=$1;
	`

	err := r.db.QueryRow(ctx, sqlQuery, userID).Scan(
		&user.Name,
		&user.Surname,
		&user.Nickname,
		&user.Email,
	)

	return &user, err
}
