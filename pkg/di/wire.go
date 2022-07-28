//go:build wireinject
// +build wireinject

package di

import (
	"go-wire-architecture/pkg/config"
	"go-wire-architecture/pkg/db"
	"go-wire-architecture/pkg/handler"
	"go-wire-architecture/pkg/handler/repository"
	http "go-wire-architecture/pkg/server"
	"go-wire-architecture/pkg/server/controllers"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(http.NewServerHTTP, db.ConnectDatabase, handler.NewUserHandler, controllers.NewUserController, controllers.NewSecondUserController, repository.NewUserRepository)

	return &http.ServerHTTP{}, nil
}
