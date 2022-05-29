package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func startWebServer(mode string) {
	if mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	registerControllers(router)

	if mode == "production" {
		log.Printf("Server is running on %s", viper.GetString("APP_URL"))
	}

	err := router.Run(viper.GetString("APP_URL"))
	if err != nil {
		log.Fatal("Failed run web server")
		return
	}
}
