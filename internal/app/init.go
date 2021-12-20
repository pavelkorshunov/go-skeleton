package app

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
