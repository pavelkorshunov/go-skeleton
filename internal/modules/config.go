package modules

import (
	"log"

	"github.com/spf13/viper"
)

func Config() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
