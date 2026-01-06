package mongo

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB(ctx context.Context) (*mongo.Client, error) {
	connString := os.Getenv("DATABASE_MONGO_URL")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connString))

	return client, err
}
