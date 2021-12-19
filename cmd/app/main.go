package main

import (
	"fmt"

	_ "github.com/pavelkorshunov/go-skeleton/internal/app"
)

var Version = "develop"

func main() {
	fmt.Println(Version)
}
