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
	timeQuery := s.db.Table("records")
	if afterCreatedAt != nil {
		query = query.Where("created_at >= ?", *afterCreatedAt)
		timeQuery = timeQuery.Where("created_at >= ?", *afterCreatedAt)
	}
	if beforeCreatedAt != nil {
		query = query.Where("created_at < ?", *beforeCreatedAt)
		timeQuery = timeQuery.Where("created_at < ?", *beforeCreatedAt)
	}

	if sensorIds != nil && len(*sensorIds) > 0 {
		query = query.Where("sensor_id in (?)", *sensorIds)
		timeQuery = timeQuery.Where("sensor_id in (?)", *sensorIds)
	}

	var times []time.Time
	var theStartTime time.Time
	var theEndTime time.Time

	rowEnd := timeQuery.Select("max(time)").Row()
	rowEnd.Scan(&theEndTime)
	log.Infof("maxTime: %v\n", theEndTime)

	row := timeQuery.Select("min(time)").Row()
	row.Scan(&theStartTime)
	log.Infof("minTime: %v\n", theStartTime)
	if startTime != nil && startTime.After(theStartTime) {
		theStartTime = *startTime
	}

	log.Infof("query minTime: %v\n", theStartTime)

	var loop int
	var base int
	if offset != nil && *offset > 0 {
		base = *offset
	} else {
		base = 0
	}
	if limit != nil && *limit > 0 {
		loop = *limit
	} else {
		loop = 10
	}
	for i := 0; i < loop; i++ {
		time := theStartTime.Add(time.Minute * time.Duration(30*i+base*30))
		times = append(times, time)
	}
	log.Infof("times: %v\n", times)

	count := theEndTime.Sub(theStartTime) / (time.Minute * time.Duration(30))
	total := int64(count)
	log.Infof("total: %v\n", total)

	err := query.Where("time in (?)", times).Statement.Order("time asc").Find(&records).Error
	if err != nil {
		log.Errorf("get records with error: %s\n", err.Error())
		return nil, err, nil
	} else {
		return records, nil, &total
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
