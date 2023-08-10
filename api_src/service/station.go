package service

import (
	"api/api_src/config"
	"api/api_src/logger"
	"api/api_src/store"
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

func (s *StationService) GetStations() ([]Station, error) {
	log := s.logger.Sugar()
	defer log.Sync()

	var stations []Station
	err := s.db.Find(&stations).Error
	if err != nil {
		log.Error(err)
		return []Station{}, err
	}
	log.Infoln("no. of stations %d retrieved", len(stations))
	return stations, nil
}
