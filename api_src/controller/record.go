package controller

import (
	"api/api_src/config"
	"api/api_src/logger"
	"api/api_src/service"
)

type RecordAPI struct {
	config        *config.Config
	logger        *logger.Logger
	recordService *service.RecordService
}

func NewRecordAPI(c *config.Config, l *logger.Logger, s *service.RecordService) *RecordAPI {
	return &RecordAPI{c, l, s}
}
