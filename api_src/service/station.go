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

func (s *StationService) GetStations() (*[]Station, error) {
	log := s.logger.Sugar()
	defer log.Sync()
	log.Infoln("get stations")
	var stations []Station
	err := s.db.Where("deleted_at IS NULL").Find(&stations).Error
	if err != nil {
		log.Error(err)
		return &[]Station{}, err
	}
	log.Infoln("no. of stations %d retrieved", len(stations))
	return &stations, nil
}

func (s *StationService) Upsert(station *Station) (*Station, error) {
	log := s.logger.Sugar()
	defer log.Sync()
	jsonVal, _ := json.Marshal(station)
	log.Infoln("upsert station: ", jsonVal)
	err := s.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "stations_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "lat", "lng", "altitude"}),
	}).Error
	return station, err
}

func (s *StationService) DeleteStation(id uint) error {
	log := s.logger.Sugar()
	defer log.Sync()
	log.Infoln("delete station with id: ", id)
	err := s.db.Model(&Station{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error
	return err
}
