package v1

import "github.com/gin-gonic/gin"

func VersionController(c *gin.Context) {
	c.JSON(200, gin.H{
		"version": "1",
	})
}
