//go:build wireinject

package api_src

import (
	"EnvMonitoringDashboard/api_src/app"
	"EnvMonitoringDashboard/api_src/config"
	"EnvMonitoringDashboard/api_src/controller"
	"EnvMonitoringDashboard/api_src/logger"
	"EnvMonitoringDashboard/api_src/service"
	"EnvMonitoringDashboard/api_src/store"
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
	service.NewSensorService,
	// api
	controller.NewNoRouteAPI,
	controller.NewDataAPI,
	controller.NewAccountAPI,
	controller.NewStationAPI,
	controller.NewRecordAPI,
	controller.NewSensorAPI,
)

func InitializeApp(webappFS fs.FS) (*app.App, error) {
	wire.Build(appSet, app.NewApp)
	return &app.App{}, nil
}
