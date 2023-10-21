package service

import (
	"EnvMonitoringDashboard/api_src/config"
	"EnvMonitoringDashboard/api_src/logger"
	"EnvMonitoringDashboard/api_src/store"
	"encoding/json"
	"time"

	"gorm.io/gorm/clause"
)

type SensorPosition string

const (
	SensorPositionUp     SensorPosition = "up"
	SensorPositionMiddle SensorPosition = "middle"
	SensorPositionDown   SensorPosition = "down"
)

var (
	SensorPositionMap = map[string]SensorPosition{
		"up":     SensorPositionUp,
		"middle": SensorPositionMiddle,
		"down":   SensorPositionDown,
	}
)

type Sensor struct {
	Base
	StationId  uint           `json:"station_id,omitempty"`
	Position   SensorPosition `json:"position,omitempty"`
	Tag        string         `json:"tag,omitempty"`
	Name       string         `json:"name,omitempty"`
	Group      string         `json:"group,omitempty"`
	Unit       string         `json:"unit,omitempty"`
	IMEI       string         `json:"imei,omitempty"`
	SensorCode string         `json:"sensor_code,omitempty"`
}

type SensorService struct {
	config *config.Config
	db     *store.DBClient
	logger *logger.Logger
}

func NewSensorService(c *config.Config, db *store.DBClient, logger *logger.Logger) *SensorService {
	log := logger.Sugar()
	defer log.Sync()

	db.AutoMigrate(&Sensor{})
	log.Infoln("database table 'sensors' migrated")

	return &SensorService{c, db, logger}
}

func (s *SensorService) Get(stationId *uint, limit *int, offset *int) (*[]Sensor, error, *int64) {
	log := s.logger.Sugar()
	defer log.Sync()
	log.Infoln("get sensors for station id: ", stationId)
	var sensors []Sensor
	var err error
	var count int64
	query := s.db.Model(&Sensor{})
	if stationId == nil {
		query = query.Where("deleted_at IS NULL")
	} else {
		query = query.Where("station_id = ? and deleted_at IS NULL", stationId)
	}
	err = query.Count(&count).Error
	if err != nil {
		log.Errorln("get sensors with error: ", err)
		return nil, err, nil
	}
	if offset != nil {
		query = query.Offset(*offset)
	}
	if limit != nil {
		query = query.Limit(*limit)
	}
	err = query.Find(&sensors).Error
	if err != nil {
		log.Errorln("get sensors with error: ", err)
		return nil, err, nil
	} else {
		return &sensors, nil, &count
	}
}

func (s *SensorService) Upsert(sensor *Sensor) (*Sensor, error) {
	log := s.logger.Sugar()
	defer log.Sync()
	jsonVal, _ := json.Marshal(sensor)
	log.Infoln("upsert sensor: ", string(jsonVal))
	err := s.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"station_id", "position", "tag", "name", "group", "unit", "updated_at", "deleted_at"}),
	}).Create(sensor).Error
	return sensor, err
}

func (s *SensorService) Delete(id uint) error {
	log := s.logger.Sugar()
	defer log.Sync()
	log.Infoln("delete sensor with id: ", id)
	err := s.db.Model(&Sensor{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error
	return err
}
