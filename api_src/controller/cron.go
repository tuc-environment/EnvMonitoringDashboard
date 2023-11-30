package controller

import (
	"EnvMonitoringDashboard/api_src/config"
	"EnvMonitoringDashboard/api_src/logger"
	"EnvMonitoringDashboard/api_src/service"
	"time"
)

type CronController struct {
	config         *config.Config
	logger         *logger.Logger
	rawDataService *service.RawDataService
	deviceService  *service.DeviceService
}

func NewCronController(c *config.Config, l *logger.Logger, s *service.RawDataService, d *service.DeviceService) *CronController {
	return &CronController{c, l, s, d}
}

func (controller *CronController) Fetch() {
	log := controller.logger.Sugar()
	defer log.Sync()
	endTime := time.Now()
	duration := -time.Minute * 30
	startTime := endTime.Add(duration)
	rawData, err := controller.rawDataService.Get(startTime, endTime)
	for _, v := range *rawData {
		controller.deviceService.ReceiveData(v.Data)
	}
	if err == nil {
		log.Infoln("Finish fetching")
	} else {
		log.Errorln("Failed feching with error: ", err)
	}
}
