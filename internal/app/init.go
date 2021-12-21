package app

import (
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
