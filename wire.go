//go:build wireinject

package main

import (
	"api/app"
	"api/config"
	"api/controller"
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
	controller.NewDataAPI,
	controller.NewAccountAPI,
	controller.NewStationAPI,
	controller.NewRecordAPI,
)

func InitializeApp() (*app.App, error) {
	wire.Build(appSet, app.NewApp)
	return &app.App{}, nil
}
