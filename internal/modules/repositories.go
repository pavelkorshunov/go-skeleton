package modules

import (
	"context"
	"github.com/pavelkorshunov/go-skeleton/internal/models"
	"github.com/pavelkorshunov/go-skeleton/internal/repository"
)

var repos *Repositories

type User interface {
	FindAll(ctx context.Context) ([]models.User, error)
}

type Repositories struct {
	User User
}

func InitRepositories() {
	db := GetConnection()

	repos = &Repositories{
		User: repository.NewUserRepo(db),
	}
}

func GetRepositories() *Repositories {
	return repos
}
