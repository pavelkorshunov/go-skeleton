package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pavelkorshunov/go-skeleton/internal/controllers/http/v1"
)

func ApiControllers(router *gin.Engine) {
	group := router.Group("/v1")
	group.GET("/version", v1.VersionController)
	group.GET("/user", v1.UserController)
}
