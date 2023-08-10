package controller

import (
	"api/config"
	"api/logger"
	"api/service"
)

type RecordAPI struct {
	config        *config.Config
	logger        *logger.Logger
	recordService *service.RecordService
}

func NewRecordAPI(c *config.Config, l *logger.Logger, s *service.RecordService) *RecordAPI {
	return &RecordAPI{c, l, s}
}
