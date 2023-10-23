package service

import (
	"EnvMonitoringDashboard/api_src/config"
	"EnvMonitoringDashboard/api_src/logger"
	"EnvMonitoringDashboard/api_src/store"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
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

type SensorSampleMethod string

const (
	SensorSampleMethodSampling SensorSampleMethod = "sampling"
	SensorSampleMethodAverage  SensorSampleMethod = "average"
	SensorSampleMethodTotal    SensorSampleMethod = "total"
	SensorSampleMethodMax      SensorSampleMethod = "max"
	SensorSampleMethodMin      SensorSampleMethod = "min"
)

var (
	SensorSampleMethodnMap = map[string]SensorSampleMethod{
		"0": SensorSampleMethodSampling,
		"1": SensorSampleMethodAverage,
		"2": SensorSampleMethodTotal,
		"3": SensorSampleMethodMax,
		"4": SensorSampleMethodMin,
	}
)

type Sensor struct {
	Base
	StationId uint           `json:"station_id,omitempty"`
	Position  SensorPosition `json:"position,omitempty"`
	Tag       string         `json:"tag,omitempty"`
	Name      string         `json:"name,omitempty"`
	Group     string         `json:"group,omitempty"`
	Unit      string         `json:"unit,omitempty"`
	// v1.1
	IMEI       string `json:"imei,omitempty"`
	SensorCode string `json:"sensor_code,omitempty"`
	// imei-sensorType-sensorCode-index
	SensorReportCode   string             `json:"sensor_report_code,omitempty"`
	SampleMethod       SensorSampleMethod `json:"sample_method,omitempty"`
	VisibleInDashboard bool               `json:"visible_in_dashboard"`
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

func (s *SensorService) Get(stationId *uint, limit *int, offset *int, visibleInDashboard *bool) (*[]Sensor, *int64, error) {
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
		log.Infof("get sensors for station id: %v\n", *stationId)
		query = query.Where("station_id = ? and deleted_at IS NULL", stationId)
	}
	if visibleInDashboard != nil {
		log.Infof("get sensors for visibleInDashboard: %v\n", *visibleInDashboard)
		query = query.Where("visible_in_dashboard = ?", *visibleInDashboard)
	}
	err = query.Count(&count).Error
	if err != nil {
		log.Errorln("get sensors with error: ", err)
		return nil, nil, err
	}
	if offset != nil {
		log.Infof("get sensors for offset: %v\n", *offset)
		query = query.Offset(*offset)
	}
	if limit != nil {
		log.Infof("get sensors for limit: %v\n", *limit)
		query = query.Limit(*limit)
	}
	err = query.Find(&sensors).Error
	if err != nil {
		log.Errorln("get sensors with error: ", err)
		return nil, nil, err
	} else {
		return &sensors, &count, nil
	}
}

func (s *SensorService) GetUnassignedSensors(limit *int, offset *int) (*[]Sensor, *int64, error) {
	log := s.logger.Sugar()
	defer log.Sync()
	log.Infoln("get unlinked sensors")
	var sensors []Sensor
	var err error
	var count int64
	query := s.db.Model(&Sensor{})
	query = query.Where("(station_id IS NULL OR station_id = 0) and deleted_at IS NULL")
	err = query.Count(&count).Error
	if err != nil {
		log.Errorln("get sensors with error: ", err)
		return nil, nil, err
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
		return nil, nil, err
	} else {
		return &sensors, &count, nil
	}
}

func (s *SensorService) GetSensorByReportCode(reportCode string) (*Sensor, error) {
	log := s.logger.Sugar()
	defer log.Sync()
	log.Infoln("get sensor by report id: ", reportCode)
	var sensor Sensor
	err := s.db.Where("sensor_report_code = ? AND deleted_at IS NULL", reportCode).First(&sensor).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err == nil {
		return &sensor, nil
	} else {
		return nil, err
	}
}

func (s *SensorService) BatchUpsert(sensors *[]Sensor) (*[]Sensor, error) {
	log := s.logger.Sugar()
	defer log.Sync()

	if sensors == nil || len(*sensors) == 0 {
		return sensors, nil
	} else {
		log.Infoln("batch upsert sensors count: ", len(*sensors))
		for i := 0; i < len(*sensors); i += 100 {
			end := i + 100
			var slice []Sensor
			if len(*sensors) < end {
				slice = (*sensors)[i:]
			} else {
				slice = (*sensors)[i:end]
			}
			err := s.db.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "id"}},
				DoUpdates: clause.AssignmentColumns([]string{"station_id", "position", "tag", "name", "group", "unit", "updated_at", "deleted_at", "imei", "sensor_code", "sensor_report_code", "sample_method", "visible_in_dashboard"}),
			}).Create(&slice).Error
			if err != nil {
				return nil, err
			}
		}
	}
	return sensors, nil
}

func (s *SensorService) Upsert(sensor *Sensor) (*Sensor, error) {
	log := s.logger.Sugar()
	defer log.Sync()
	jsonVal, _ := json.Marshal(sensor)
	log.Infoln("upsert sensor: ", string(jsonVal))
	err := s.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"station_id", "position", "tag", "name", "group", "unit", "updated_at", "deleted_at", "imei", "sensor_code", "sensor_report_code", "sample_method", "visible_in_dashboard"}),
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
