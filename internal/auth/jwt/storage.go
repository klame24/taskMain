package jwt

import (
	"context"
	"crypto/rand"
	"encoding/hex"
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "expires_at", Value: 1}},
		Options: options.Index().SetExpireAfterSeconds(0),
	}

	collection.Indexes().CreateOne(ctx, indexModel)

	return &TokenRepository{collection: collection}
}

func (r *TokenRepository) SaveRefreshToken(ctx context.Context, userID int, token string, expiresAt time.Time) error {
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

func GenerateRefreshToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
