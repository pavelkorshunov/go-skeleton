package modules

import (
	"context"
	"log"

	"github.com/pavelkorshunov/go-skeleton/pkg/mongodb"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

var connection *mongo.Database

func SetConnection(ctx context.Context, db string) {
	client, err := mongodb.NewClient(ctx, viper.GetString("MONGODB_URI"))
	if err != nil {
		log.Fatal("Failed to connect mongodb")
	}

	connection = client.Database(db)
}

func GetConnection() *mongo.Database {
	return connection
}
