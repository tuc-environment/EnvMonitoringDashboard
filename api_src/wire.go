//go:build wireinject

package api_src

import (
	"api/api_src/app"
	"api/api_src/config"
	"api/api_src/controller"
	"api/api_src/logger"
	"api/api_src/service"
	"api/api_src/store"
	"io/fs"

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
	controller.NewNoRouteAPI,
	controller.NewDataAPI,
	controller.NewAccountAPI,
	controller.NewStationAPI,
	controller.NewRecordAPI,
)

func InitializeApp(webappFS fs.FS) (*app.App, error) {
	wire.Build(appSet, app.NewApp)
	return &app.App{}, nil
}
