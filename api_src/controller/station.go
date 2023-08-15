package controller

import (
	"EnvMonitoringDashboard/api_src/config"
	"EnvMonitoringDashboard/api_src/controller/args"
	"EnvMonitoringDashboard/api_src/logger"
	"EnvMonitoringDashboard/api_src/service"

	"github.com/gin-gonic/gin"
)

type StationAPI struct {
	config         *config.Config
	logger         *logger.Logger
	stationService *service.StationService
}

func NewStationAPI(c *config.Config, l *logger.Logger, s *service.StationService) *StationAPI {
	return &StationAPI{c, l, s}
}

// Get stations godoc
//
//	@Summary		Get stations
//	@Description	Get stations with params
//	@Tags			stations
//	@Accept			json
//	@Produce		json
//	@Success		200	"Return station json array"
//	@Router			/stations [get]
func (api *StationAPI) GetStations(g *gin.Context) {
	log := api.logger.Sugar()
	defer log.Sync()
	c := WrapContext(g)
	stations, err := api.stationService.GetStations()
	if err != nil {
		c.BadRequest(err)
		return
	}
	c.OK(stations)
}

// Upsert stations godoc
//
//	@Summary		Upsert stations
//	@Description	Upsert stations
//	@Tags			stations
//	@Accept			json
//	@Produce		json
//	@Success		200	"Return station json array"
//	@Router			/stations [post]
func (api *StationAPI) UpsertStations(g *gin.Context) {
	log := api.logger.Sugar()
	defer log.Sync()
	c := WrapContext(g)
	body := args.UpsertStationArgs{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.BadRequest(err)
		return
	}
	station := service.Station{
		Base: service.Base{
			ID: body.ID,
		},
		Name:     body.Name,
		Lat:      body.Lat,
		Lng:      body.Lng,
		Altitude: body.Altitude,
	}
	api.stationService.Upsert(&station)
	c.OK(station)
}
