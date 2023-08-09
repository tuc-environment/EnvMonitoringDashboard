package service

import (
	"api/config"
	"api/logger"
	"api/store"
	"time"
)

type RecordPosition string

const (
	RecordPositionBelow  RecordPosition = "below"
	RecordPositionMiddle RecordPosition = "middle"
	RecordPositionAbove  RecordPosition = "above"
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
