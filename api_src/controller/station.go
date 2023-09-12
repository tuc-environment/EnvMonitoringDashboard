package controller

import (
	"EnvMonitoringDashboard/api_src/config"
	"EnvMonitoringDashboard/api_src/controller/args"
	"EnvMonitoringDashboard/api_src/logger"
	"EnvMonitoringDashboard/api_src/service"
	"EnvMonitoringDashboard/api_src/utils"
	"errors"
	"strconv"

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
	q := c.Request.URL.Query()
	var offset *int
	var limit *int
	if q.Has("offset") {
		str := q.Get("offset")
		offsetV, _ := strconv.Atoi(str)
		offset = &offsetV
	}
	if q.Has("limit") {
		str := q.Get("limit")
		limitV, _ := strconv.Atoi(str)
		limit = &limitV
	}
	stations, count, err := api.stationService.GetStations(offset, limit)
	if err != nil {
		c.BadRequest(err)
		return
	}
	c.Header(utils.XTotalCount, strconv.FormatInt(*count, 10))
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

// Delete stations godoc
//
//	@Summary		delete stations
//	@Description	delete stations
//	@Tags			stations
//	@Accept			json
//	@Produce		json
//	@Success		200	"Return station id"
//	@Router			/stations/:station_id [delete]
func (api *StationAPI) DeleteStation(g *gin.Context) {
	log := api.logger.Sugar()
	defer log.Sync()
	c := WrapContext(g)
	stationId := c.Param("station_id")
	num, err := strconv.Atoi(stationId)
	if err != nil {
		c.BadRequest(errors.New("invalid station_id"))
		return
	}
	err = api.stationService.DeleteStation(uint(num))
	if err != nil {
		c.BadRequest(err)
	} else {
		c.OK(nil)
	}
}
