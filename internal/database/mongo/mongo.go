package mongo

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectMongoDB(ctx context.Context) (*mongo.Client, error) {
	connString := os.Getenv("DATABASE_MONGO_URL")

	client, err := mongo.Connect(options.Client().ApplyURI(connString))

	return client, err
}
