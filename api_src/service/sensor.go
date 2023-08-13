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
	StationId uint         `json:"station_id,omitempty"`
	Position  SensorPosition `json:"position,omitempty"`
	Tag       string         `json:"tag,omitempty"`
	Name      string         `json:"name,omitempty"`
	Group     string         `json:"group,omitempty"`
	Unit      string         `json:"unit,omitempty"`
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

func (s *SensorService) GetSensor(stationId uint) (*[]Sensor, error) {
	log := s.logger.Sugar()
	defer log.Sync()
	log.Infoln("get sensors for station id: ", stationId)
	var sensors *[]Sensor
	err := s.db.Find(sensors).Where("station_id = ? and deleted_at IS NOT NULL", stationId).Error
	if err != nil {
		log.Errorln("get sensors with error: ", err)
		return nil, err
	} else {
		return sensors, nil
	}
}

func (s *SensorService) Upsert(sensor *Sensor) (*Sensor, error) {
	log := s.logger.Sugar()
	defer log.Sync()
	jsonVal, _ := json.Marshal(sensor)
	log.Infoln("upsert sensor: ", jsonVal)
	err := s.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"position", "tag", "name", "group", "unit", "updated_at", "deleted_at"}),
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
