package mongodb

import (
	"context"
	"github.com/EDEN-NN/hydra-api/internal/apperrors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func Connect() (*mongo.Database, error) {
	uri := "mongodb://root:example@localhost:27017/hydradb/auth?authSource=admin"

	log.Println("Connecting to MongoDB...")

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(apperrors.NewError(apperrors.EINTERNAL, "fail to connect to mongodb", err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5+time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB")
		return nil, apperrors.NewError(apperrors.EINTERNAL, "fail to comunicate to mongo", err)
	}

	log.Println("Successfully connected to MongoDB!")

	return client.Database("hydradb"), nil
}

func ConnectTest() (*mongo.Database, error) {
	uri := os.Getenv("MONGO_URI")

	log.Println("Connecting to MongoDB...")

	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerSelectionTimeout(10 * time.Second).
		SetConnectTimeout(10 * time.Second)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(apperrors.NewError(apperrors.EINTERNAL, "fail to connect to mongodb", err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5+time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB")
		return nil, apperrors.NewError(apperrors.EINTERNAL, "fail to comunicate to mongo", err)
	}

	log.Println("Successfully connected to MongoDB!")

	return client.Database("hydradb_test"), nil
}
