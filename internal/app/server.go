package app

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func startWebServer(mode string) {
	if mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	registerControllers(router)
	router.Run(viper.GetString("APP_URL"))
}
