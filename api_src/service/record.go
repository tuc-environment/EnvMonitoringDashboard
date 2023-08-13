package service

import (
	"EnvMonitoringDashboard/api_src/config"
	"EnvMonitoringDashboard/api_src/logger"
	"EnvMonitoringDashboard/api_src/store"
	"time"

	"gorm.io/gorm/clause"
)

type Record struct {
	Base
	SensorId uint    `json:"sensor_id,omitempty" gorm:"uniqueIndex:sensor_time"`
	Value    float64   `json:"value,omitempty"`
	Time     time.Time `json:"time,omitempty" gorm:"uniqueIndex:sensor_time"`
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

func (s *RecordService) GetRecords(sensorId uint) (*[]Record, error) {
	log := s.logger.Sugar()
	defer log.Sync()
	var records *[]Record
	err := s.db.Find(records).Where("sensor_id = ? and deleted_at IS NOT NULL", sensorId).Error
	if err != nil {
		log.Errorf("get records for sensor_id: %d, error: %s\n", sensorId, err.Error())
		return nil, err
	} else {
		return records, nil
	}
}

func (s *RecordService) BatchUpsert(records *[]Record) (*[]Record, error) {
	log := s.logger.Sugar()
	defer log.Sync()
	log.Infoln("batch upsert records count: ", len(*records))
	err := s.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "sensor_id"}, {Name: "time"}},
		DoUpdates: clause.AssignmentColumns([]string{"value", "updated_at", "deleted_at"}),
	}).Create(records).Error
	return records, err
}

func (s *RecordService) DeleteRecord(id uint) error {
	log := s.logger.Sugar()
	defer log.Sync()
	log.Infoln("delete record with id: ", id)
	err := s.db.Model(&Record{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error
	return err
}
