//go:build wireinject

package main

import (
	"api/api"
	"api/app"
	"api/config"
	"api/db"
	"api/logger"

	"github.com/google/wire"
)

var appSet = wire.NewSet(
	app.NewEngine,
	// config and logger
	config.NewConfig,
	logger.NewLogger,
	// store
	db.NewDB,
	db.NewService,
	// api
	api.NewHelloAPI,
)

func InitializeApp() (*app.App, error) {
	wire.Build(appSet, app.NewApp)
	return &app.App{}, nil
}
