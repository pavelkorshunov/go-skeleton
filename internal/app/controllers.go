package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pavelkorshunov/go-skeleton/internal/routes"
)

func registerControllers(router *gin.Engine) {
	routes.ApiControllers(router)
}
