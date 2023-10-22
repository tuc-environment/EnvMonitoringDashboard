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
	SensorId    uint      `json:"sensor_id,omitempty" gorm:"uniqueIndex:sensor_time"`
	Value       float64   `json:"value,omitempty"`
	Time        time.Time `json:"time,omitempty" gorm:"uniqueIndex:sensor_time"`
	RecordIndex uint64    `json:"record_id,omitempty"`
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
	var records []Record

	query := s.db.DB.Model(&Record{})
	// timeQuery := s.db.Table("records"
	timeQuery := s.db.DB.Model(&Record{})
	startTimeQuery := s.db.DB.Model(&Record{})
	if afterCreatedAt != nil {
		query = query.Where("created_at >= ?", *afterCreatedAt)
		timeQuery = timeQuery.Where("created_at >= ?", *afterCreatedAt)
		startTimeQuery = startTimeQuery.Where("created_at >= ?", *afterCreatedAt)
	}
	if beforeCreatedAt != nil {
		query = query.Where("created_at < ?", *beforeCreatedAt)
		timeQuery = timeQuery.Where("created_at < ?", *beforeCreatedAt)
		startTimeQuery = startTimeQuery.Where("created_at < ?", *beforeCreatedAt)
	}

	if sensorIds != nil && len(*sensorIds) > 0 {
		query = query.Where("sensor_id in (?)", *sensorIds)
		timeQuery = timeQuery.Where("sensor_id in (?)", *sensorIds)
		startTimeQuery = startTimeQuery.Where("sensor_id in (?)", *sensorIds)
	}

	var theStartTime time.Time
	var startTimeRecords *[]Record
	var theEndTime time.Time
	var endTimeRecords *[]Record

	endQueryErr := timeQuery.Limit(1).Order("time desc").Find(&endTimeRecords).Error
	if endQueryErr == nil && endTimeRecords != nil && len(*endTimeRecords) > 0 {
		theEndTime = (*endTimeRecords)[0].Time
		log.Infof("maxTime: %v\n", theEndTime)
	}

	startTimeErr := startTimeQuery.Limit(1).Order("time asc").Find(&startTimeRecords).Error
	if startTimeErr == nil && startTimeRecords != nil && len(*startTimeRecords) > 0 {
		theStartTime = (*startTimeRecords)[0].Time
		log.Infof("minTime: %v\n", theStartTime)
	}

	if startTime != nil && startTime.After(theStartTime) {
		theStartTime = *startTime
	}

	log.Infof("query minTime: %v\n", theStartTime)

	var queryStartTime time.Time
	var queryEndTime time.Time
	if offset != nil && *offset > 0 {
		queryStartTime = theStartTime.Add(time.Minute * time.Duration(30*(*offset)))
	} else {
		queryStartTime = theStartTime
	}
	if limit != nil && *limit > 0 {
		queryEndTime = queryStartTime.Add(time.Minute * time.Duration(30*(*limit)))
	} else {
		queryEndTime = queryStartTime.Add(time.Minute * time.Duration(300))
	}

	log.Infof("query start time: %v, end time: %v\n", queryStartTime, queryEndTime)

	count := theEndTime.Sub(theStartTime) / (time.Minute * time.Duration(30))
	total := int64(count)
	log.Infof("total: %v\n", total)

	err := query.Where("time >= ?", queryStartTime).Where("time < ?", queryEndTime).Order("time asc").Find(&records).Error
	if err != nil {
		log.Errorf("get records with error: %s\n", err.Error())
		return nil, err, nil
	} else {
		return &records, nil, &total
	}
}

func (s *RecordService) BatchUpsert(records *[]Record) (*[]Record, error) {
	log := s.logger.Sugar()
	defer log.Sync()

	if records == nil || len(*records) == 0 {
		return records, nil
	} else {
		log.Infoln("batch upsert records count: ", len(*records))
		for i := 0; i < len(*records); i += 100 {
			end := i + 100
			var slice []Record
			if len(*records) < end {
				slice = (*records)[i:]
			} else {
				slice = (*records)[i:end]
			}
			err := s.db.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "sensor_id"}, {Name: "time"}},
				DoUpdates: clause.AssignmentColumns([]string{"value", "updated_at", "deleted_at"}),
			}).Create(&slice).Error
			if err != nil {
				return nil, err
			}
		}
	}
	return records, nil
}

func (s *RecordService) DeleteRecord(id uint) error {
	log := s.logger.Sugar()
	defer log.Sync()
	log.Infoln("delete record with id: ", id)
	err := s.db.Model(&Record{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error
	return err
}
