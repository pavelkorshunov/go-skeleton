package modules

import (
	"log"

	"github.com/pavelkorshunov/go-skeleton/pkg/mongodb"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

var connection *mongo.Database

func ConnectionDatabase() {
	uri := viper.GetString("DB_URI")
	username := viper.GetString("DB_USERNAME")
	password := viper.GetString("DB_PASSWORD")

	client, err := mongodb.NewClient(uri, username, password)
	if err != nil {
		log.Fatal("Failed to connect mongodb")
	}

	connection = client.Database(viper.GetString("DB_DATABASE"))
}

func GetConnection() *mongo.Database {
	return connection
}
