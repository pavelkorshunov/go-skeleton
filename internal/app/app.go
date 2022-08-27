package app

import (
	"github.com/pavelkorshunov/go-skeleton/internal/modules"
)

var Mode = "develop"

func Run() {
	modules.SetConnection()

	startWebServer(Mode)
}
