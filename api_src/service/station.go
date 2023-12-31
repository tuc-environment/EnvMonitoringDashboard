package service

import (
	"EnvMonitoringDashboard/api_src/config"
	"EnvMonitoringDashboard/api_src/logger"
	"EnvMonitoringDashboard/api_src/store"
	"encoding/json"
	"time"

	"gorm.io/gorm/clause"
)

type Station struct {
	Base
	Name     string  `json:"name"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
	Altitude float64 `json:"altitude"`
}

type StationPredictionResult struct {
	DownSoilTempShallow           float64 `json:"down_soil_temp_shallow"`
	DownSoilTempDeep              float64 `json:"down_soil_temp_deep"`
	DownSoilWaterContentShallow   float64 `json:"down_soil_water_content_shallow"`
	DownSoilWaterContentDeep      float64 `json:"down_soil_water_content_deep"`
	DownTemp                      float64 `json:"down_temp"`
	DownHumidity                  float64 `json:"down_humidity"`
	MiddleSoilTempShallow         float64 `json:"middle_soil_temp_shallow"`
	MiddleSoilTempDeep            float64 `json:"middle_soil_temp_deep"`
	MiddleSoilWaterContentShallow float64 `json:"middle_soil_water_content_shallow"`
	MiddleSoilWaterContentDeep    float64 `json:"middle_soil_water_content_deep"`
	MiddleTemp                    float64 `json:"middle_temp"`
	MiddleHumidity                float64 `json:"middle_humidity"`
}

type StationService struct {
	config *config.Config
	db     *store.DBClient
	logger *logger.Logger
}

func NewStationService(c *config.Config, db *store.DBClient, logger *logger.Logger) *StationService {
	log := logger.Sugar()
	defer log.Sync()

	db.AutoMigrate(&Station{})
	log.Infoln("database table 'stations' migrated")

	return &StationService{c, db, logger}
}

func (s *StationService) GetStations(offset *int, limit *int) (*[]Station, *int64, error) {
	log := s.logger.Sugar()
	defer log.Sync()
	log.Infoln("get stations")
	var stations []Station
	var count int64
	query := s.db.Model(&Station{}).Where("deleted_at IS NULL")
	err := query.Count(&count).Error
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}
	if offset != nil {
		query = query.Offset(*offset)
	}
	if limit != nil {
		query = query.Limit(*limit)
	}
	err = query.Find(&stations).Error
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}
	log.Infoln("no. of stations %d retrieved", len(stations))
	return &stations, &count, nil
}

func (s *StationService) Upsert(station *Station) (*Station, error) {
	log := s.logger.Sugar()
	defer log.Sync()
	jsonVal, _ := json.Marshal(station)
	log.Infoln("upsert station: ", string(jsonVal))
	err := s.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "lat", "lng", "altitude"}),
	}).Create(station).Error
	return station, err
}

func (s *StationService) DeleteStation(id uint) error {
	log := s.logger.Sugar()
	defer log.Sync()
	log.Infoln("delete station with id: ", id)
	err := s.db.Model(&Station{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error
	return err
}
