package service

import (
	"EnvMonitoringDashboard/api_src/config"
	"EnvMonitoringDashboard/api_src/logger"
	"EnvMonitoringDashboard/api_src/store"
	"time"
)

type RawData struct {
	IMEI         string    `gorm:"column:imei"`
	Time         time.Time `gorm:"column:tmstamp"`
	Index        int       `gorm:"column:recnum"`
	Data         string    `gorm:"column:data"`
	ReceivedTime time.Time `gorm:"column:recvtime"`
}

func (RawData) TableName() string {
	return "records"
}

type RawDataService struct {
	config *config.Config
	db     *store.RawDBClient
	logger *logger.Logger
}

func NewRawDataService(c *config.Config, db *store.RawDBClient, logger *logger.Logger) *RawDataService {
	log := logger.Sugar()
	defer log.Sync()
	return &RawDataService{c, db, logger}
}

func (s *RawDataService) Get(startTime time.Time, endTime time.Time) (*[]RawData, error) {
	log := s.logger.Sugar()
	defer log.Sync()
	var rawRecords []RawData
	err := s.db.DB.Model(&RawData{}).Where("tmstamp >= ?", startTime).Where("tmstamp < ?", endTime).Order("tmstamp asc").Find(&rawRecords).Error
	if err != nil {
		log.Errorf("get raw data with error: %s\n", err.Error())
		return nil, err
	} else {
		return &rawRecords, nil
	}
}
