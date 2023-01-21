package app

import (
	"github.com/pavelkorshunov/go-skeleton/internal/modules"
)

func init() {
	modules.Config()
	modules.ConnectionDatabase()
	modules.InitRepositories()
}
