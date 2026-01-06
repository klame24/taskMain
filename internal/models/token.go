package models

import "time"

type RefreshToken struct {
	UserID    int       `bson:"user_id"`
	TokenHash string    `bson:"token_hash"`
	ExpiresAt time.Time `bson:"expires_at"`
	CreatedAt time.Time `bson:"created_at"`
}
