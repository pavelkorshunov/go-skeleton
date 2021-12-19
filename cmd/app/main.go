package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/pavelkorshunov/go-skeleton/internal/app"
	v1 "github.com/pavelkorshunov/go-skeleton/internal/controllers/http/v1"
)

var Version = "develop"

func main() {
	router := gin.Default()
	router.GET("/", v1.HelloController)
	router.Run(":8080")
}
