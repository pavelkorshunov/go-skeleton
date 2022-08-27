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

	addr := viper.GetString("APP_URL")

	router := gin.Default()
	registerControllers(router)

	if mode == "production" {
		log.Printf("Server is running on %s", addr)
	}

	err := router.Run(addr)
	if err != nil {
		panic("Failed run web server")
	}
}
