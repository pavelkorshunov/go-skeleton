package app

import (
	"context"
	"time"

	"github.com/pavelkorshunov/go-skeleton/internal/modules"
)

var Mode = "develop"

func Run() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	modules.SetConnection(ctx, "testing")

	startWebServer(Mode)
}
