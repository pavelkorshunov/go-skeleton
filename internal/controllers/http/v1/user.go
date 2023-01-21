package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pavelkorshunov/go-skeleton/internal/modules"
)

func UserController(c *gin.Context) {
	repos := modules.GetRepositories()
	users, _ := repos.User.FindAll(c.Request.Context())

	c.JSON(http.StatusOK, users)
}
