//go:build wireinject

package main

import (
	"api/api"
	"api/app"
	"api/config"

	"github.com/google/wire"
)

var appSet = wire.NewSet(
	app.NewEngine,
	config.NewConfig,
	api.NewHelloAPI,
)

func InitializeApp() (*app.App, error) {
	wire.Build(appSet, app.NewApp)
	return &app.App{}, nil
}
