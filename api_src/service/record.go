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
	SensorId uint      `json:"sensor_id,omitempty" gorm:"uniqueIndex:sensor_time"`
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

func (s *RecordService) GetRecords(sensorIds *[]uint, startTime *time.Time, endTime *time.Time, afterCreatedAt *time.Time, beforeCreatedAt *time.Time, offset *int, limit *int) (*[]Record, error, *int64) {
	log := s.logger.Sugar()
	defer log.Sync()
	var records *[]Record
	query := s.db.DB.Model(&Record{})
	if afterCreatedAt != nil {
		query = query.Where("created_at >= ?", *afterCreatedAt)
	}
	if beforeCreatedAt != nil {
		query = query.Where("created_at < ?", *beforeCreatedAt)
	}
	if startTime != nil {
		query = query.Where("time >= ?", *startTime)
	}
	if endTime != nil {
		query = query.Where("time < ?", *endTime)
	}
	if sensorIds != nil && len(*sensorIds) > 0 {
		query = query.Where("sensor_id in (?)", *sensorIds)
	}
	var count int64
	errCount := query.Count(&count).Error
	if errCount != nil {
		return nil, errCount, nil
	}
	if offset != nil {
		query = query.Offset(*offset)
	}
	if limit != nil {
		query = query.Limit(*limit)
	}
	err := query.Statement.Order("time desc").Find(&records).Error
	if err != nil {
		log.Errorf("get records with error: %s\n", err.Error())
		return nil, err, nil
	} else {
		return records, nil, &count
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
