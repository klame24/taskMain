package jwt

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"taskMain/internal/auth/password"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TokenRepository struct {
	collection *mongo.Collection
}

func NewTokenRepository(client *mongo.Client) *TokenRepository {
	db := client.Database("auth_db")
	collection := db.Collection("refresh_tokens")

	// Создаем TTL индекс для автоматического удаления просроченных токенов
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"expires_at", 1}},
		Options: options.Index().SetExpireAfterSeconds(0),
	}

	collection.Indexes().CreateOne(ctx, indexModel)

	return &TokenRepository{collection: collection}
}

func (r *TokenRepository) SaveRefreshToken(ctx context.Context, userID int, token string, expiresAt time.Time) error {
	// Хэшируем refresh token перед сохранением
	tokenHash, err := password.HashPassword(token)
	if err != nil {
		return err
	}

	_, err = r.collection.InsertOne(ctx, bson.M{
		"user_id":    userID,
		"token_hash": tokenHash,
		"expires_at": expiresAt,
		"created_at": time.Now(),
	})

	return err
}

func (r *TokenRepository) GetUserByRefreshToken(ctx context.Context, token string) (int, error) {
	// Находим токен по хэшу
	// var result struct {
	// 	UserID int `bson:"user_id"`
	// }

	// Нужно перебрать все токены и проверить каждый (bcrypt.Compare)
	// Для простоты пока что не реализуем полностью
	return 0, fmt.Errorf("not implemented")
}

func (r *TokenRepository) DeleteRefreshToken(ctx context.Context, token string) error {
	// Аналогично, нужно найти по хэшу
	return fmt.Errorf("not implemented")
}

func (r *TokenRepository) DeleteAllForUser(ctx context.Context, userID int) error {
	_, err := r.collection.DeleteMany(ctx, bson.M{"user_id": userID})
	return err
}

// Вспомогательная функция для генерации refresh token
func GenerateRefreshToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
