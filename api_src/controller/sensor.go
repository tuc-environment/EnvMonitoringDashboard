package controller

import (
	"EnvMonitoringDashboard/api_src/config"
	"EnvMonitoringDashboard/api_src/logger"
	"EnvMonitoringDashboard/api_src/service"
	"EnvMonitoringDashboard/api_src/utils"
	"errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type SensorAPI struct {
	config        *config.Config
	logger        *logger.Logger
	sensorService *service.SensorService
}

func NewSensorAPI(c *config.Config, l *logger.Logger, s *service.SensorService) *SensorAPI {
	return &SensorAPI{c, l, s}
}

// Get sensors godoc
//
//	@Summary	get sensors
//	@Tags		sensors
//	@Accept		json
//	@Produce	json
//	@Success	200	"return sensors by station_id"
//	@Router		/sensors [get]
func (api *SensorAPI) GetSensors(g *gin.Context) {
	log := api.logger.Sugar()
	defer log.Sync()
	c := WrapContext(g)
	var stationId *uint
	var offset *int
	var limit *int
	var visibleInDashboard *bool
	q := c.Request.URL.Query()
	if q.Has("station_id") {
		str := q.Get("station_id")
		num, err := strconv.Atoi(str)
		if err != nil {
			c.BadRequest(errors.New("invalid station_id"))
			return
		}
		unitV := uint(num)
		stationId = &unitV
	}
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
	if q.Has("visible_in_dashboard") {
		str := q.Get("visible_in_dashboard")
		visibleInDashboardV, _ := strconv.ParseBool(str)
		visibleInDashboard = &visibleInDashboardV
	}
	sensors, count, err := api.sensorService.Get(stationId, offset, limit, visibleInDashboard)
	if err != nil {
		c.BadRequest(err)
	} else {
		c.Header(utils.XTotalCount, strconv.FormatInt(*count, 10))
		c.OK(sensors)
	}
}

// Get sensors godoc
//
//	@Summary	get sensors
//	@Tags		sensors
//	@Accept		json
//	@Produce	json
//	@Success	200	"return sensors by station_id"
//	@Router		/sensors [get]
func (api *SensorAPI) GetUnassignedSensors(g *gin.Context) {
	log := api.logger.Sugar()
	defer log.Sync()
	c := WrapContext(g)
	var offset *int
	var limit *int
	q := c.Request.URL.Query()
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
	sensors, count, err := api.sensorService.GetUnassignedSensors(offset, limit)
	if err != nil {
		c.BadRequest(err)
	} else {
		c.Header(utils.XTotalCount, strconv.FormatInt(*count, 10))
		c.OK(sensors)
	}
}

// Upsert sensor godoc
//
//	@Summary	create sensor
//	@Tags		sensors
//	@Accept		json
//	@Produce	json
//	@Success	200	"return sensor"
//	@Router		/sensors [post]
func (api *SensorAPI) UpsertSensor(g *gin.Context) {
	log := api.logger.Sugar()
	defer log.Sync()
	c := WrapContext(g)
	var sensor *service.Sensor
	sensor = &service.Sensor{}
	if err := c.ShouldBindJSON(&sensor); err != nil {
		c.BadRequest(err)
		return
	}
	sensor, err := api.sensorService.Upsert(sensor)
	if err != nil {
		c.BadRequest(err)
	} else {
		c.OK(sensor)
	}
}

// Batch upsert sensor godoc
//
//	@Summary		batch upsert
//	@Description	batch upsert sensors
//	@Tags			sensors
//	@Accept			json
//	@Produce		json
//	@Success		200	""
//	@Router			/sensors [post]
func (api *SensorAPI) BatchUpsertSensors(g *gin.Context) {
	log := api.logger.Sugar()
	defer log.Sync()
	c := WrapContext(g)
	body := []service.Sensor{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.BadRequest(err)
		return
	}
	records, err := api.sensorService.BatchUpsert(&body)
	if err != nil {
		c.BadRequest(err)
	} else {
		c.OK(records)
	}
}

// Delete sensors godoc
//
//	@Summary		delete sensors
//	@Description	delete sensors
//	@Tags			sensors
//	@Accept			json
//	@Produce		json
//	@Success		200	"Return sensor id"
//	@Router			/sensors/:sensor_id [delete]
func (api *SensorAPI) DeleteSensor(g *gin.Context) {
	log := api.logger.Sugar()
	defer log.Sync()
	c := WrapContext(g)
	sensorId := c.Param("sensor_id")
	num, err := strconv.Atoi(sensorId)
	if err != nil {
		c.BadRequest(errors.New("invalid sensor_id"))
		return
	}
	err = api.sensorService.Delete(uint(num))
	if err != nil {
		c.BadRequest(err)
	} else {
		c.OK(nil)
	}
}

func (api *SensorAPI) parsePosition(str string) (service.SensorPosition, bool) {
	c, ok := service.SensorPositionMap[strings.ToLower(str)]
	return c, ok
}
