package service

import (
	"EnvMonitoringDashboard/api_src/config"
	"EnvMonitoringDashboard/api_src/logger"
	"EnvMonitoringDashboard/api_src/store"
	"time"
)

type RecordPosition string

const (
	RecordPositionUp     RecordPosition = "up"
	RecordPositionMiddle RecordPosition = "middle"
	RecordPositionDown   RecordPosition = "down"
)

var (
	RecordPositionMap = map[string]RecordPosition{
		"up":     RecordPositionUp,
		"middle": RecordPositionMiddle,
		"down":   RecordPositionDown,
	}
)

type Record struct {
	Base
	StationId string         `json:"station_id"`
	Position  RecordPosition `json:"position"`
	Tag       string         `json:"tag"`
	Name      string         `json:"name"`
	Sensor    string         `json:"sensor"`
	Group     string         `json:"group"`
	Value     float64        `json:"value"`
	Unit      string         `json:"unit"`
	Time      time.Time      `json:"time"`
}

type RecordService struct {
	config *config.Config
	db     *store.DBClient
	logger *logger.Logger
}

func NewRecordService(c *config.Config, db *store.DBClient, logger *logger.Logger) *RecordService {
	log := logger.Sugar()
	defer log.Sync()

	db.AutoMigrate(&Record{})
	log.Infoln("database table 'records' migrated")

	return &RecordService{c, db, logger}
}

func (s *RecordService) UploadRecords() ([]Record, error) {
	log := s.logger.Sugar()
	defer log.Sync()
	var records []Record

	log.Infoln("no. of records %d retrieved", len(records))
	return nil, nil
}
