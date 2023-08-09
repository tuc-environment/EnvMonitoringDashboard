//go:build wireinject

package main

import (
	"api/api"
	"api/app"
	"api/config"
	"api/logger"
	"api/service"
	"api/store"

	"github.com/google/wire"
)

var appSet = wire.NewSet(
	app.NewEngine,
	// config and logger
	config.NewConfig,
	logger.NewLogger,
	// store
	store.NewDBClient,
	// service
	service.NewAccountService,
	service.NewStationService,
	service.NewRecordService,
	// api
	api.NewDataAPI,
	api.NewAccountAPI,
	api.NewStationAPI,
	api.NewRecordAPI,
)

func InitializeApp() (*app.App, error) {
	wire.Build(appSet, app.NewApp)
	return &app.App{}, nil
}
