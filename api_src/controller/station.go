package controller

import (
	"EnvMonitoringDashboard/api_src/config"
	"EnvMonitoringDashboard/api_src/controller/args"
	"EnvMonitoringDashboard/api_src/logger"
	"EnvMonitoringDashboard/api_src/service"
	"EnvMonitoringDashboard/api_src/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
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

// Predict stations godoc
//
//	@Summary		predict station data
//	@Description	predict station data
//	@Tags			stations
//	@Accept			json
//	@Produce		json
//	@Success		200	"Return station prediction data"
//	@Router			/stations/predict [delete]
func (api *StationAPI) PredictStation(g *gin.Context) {
	log := api.logger.Sugar()
	defer log.Sync()
	c := WrapContext(g)
	q := c.Request.URL.Query()
	var lat *float64
	var lng *float64

	var temp *float64
	var humidity *float64
	var barometricPressure *float64
	var soilTempShallow *float64
	var soilTempDeep *float64
	var soilWaterContentShallow *float64
	var soilWaterContentDeep *float64
	var soilElectricalConductivity *float64

	if q.Has("lat") {
		str := q.Get("lat")
		latVal, err := strconv.ParseFloat(str, 64)
		if err == nil {
			lat = &latVal
		}
	}
	if q.Has("lng") {
		str := q.Get("lng")
		lngVal, err := strconv.ParseFloat(str, 64)
		if err == nil {
			lng = &lngVal
		}
	}

	if q.Has("temp") {
		str := q.Get("temp")
		tempVal, err := strconv.ParseFloat(str, 64)
		if err == nil {
			temp = &tempVal
		}
	}
	if q.Has("humidity") {
		str := q.Get("humidity")
		humidityVal, err := strconv.ParseFloat(str, 64)
		if err == nil {
			humidity = &humidityVal
		}
	}
	if q.Has("barometric_pressure") {
		str := q.Get("barometric_pressure")
		barometricPressureVal, err := strconv.ParseFloat(str, 64)
		if err == nil {
			barometricPressure = &barometricPressureVal
		}
	}
	if q.Has("soil_temp_shallow") {
		str := q.Get("soil_temp_shallow")
		soilTempShallowVal, err := strconv.ParseFloat(str, 64)
		if err == nil {
			soilTempShallow = &soilTempShallowVal
		}
	}
	if q.Has("soil_temp_deep") {
		str := q.Get("soil_temp_deep")
		soilTempDeepVal, err := strconv.ParseFloat(str, 64)
		if err == nil {
			soilTempDeep = &soilTempDeepVal
		}
	}
	if q.Has("soil_water_content_shallow") {
		str := q.Get("soil_water_content_shallow")
		soilWaterContentShallowVal, err := strconv.ParseFloat(str, 64)
		if err == nil {
			soilWaterContentShallow = &soilWaterContentShallowVal
		}
	}
	if q.Has("soil_water_content_deep") {
		str := q.Get("soil_water_content_deep")
		soilWaterContentDeepVal, err := strconv.ParseFloat(str, 64)
		if err == nil {
			soilWaterContentDeep = &soilWaterContentDeepVal
		}
	}
	if q.Has("soil_electrical_conductivity") {
		str := q.Get("soil_electrical_conductivity")
		soilElectricalConductivityVal, err := strconv.ParseFloat(str, 64)
		if err == nil {
			soilElectricalConductivity = &soilElectricalConductivityVal
		}
	}
	log.Infof("[station-prediction] temp: %v, humidity: %v, barometricPressure: %v, soilTempShallow: %v, soilTempDeep: %v, soilWaterContentShallow: %v, soilWaterContentDeep: %v, soilElectricalConductivity: %v", temp, humidity, barometricPressure, soilTempShallow, soilTempDeep, soilWaterContentShallow, soilWaterContentDeep, soilElectricalConductivity)
	if temp != nil && humidity != nil && barometricPressure != nil && soilTempShallow != nil && soilTempDeep != nil && soilWaterContentShallow != nil && soilWaterContentDeep != nil && soilElectricalConductivity != nil {
		if lat != nil && lng != nil {
			log.Infof("[station-prediction] lat: %v, lng: %v", *lat, *lng)
		}
		log.Infof("[station-prediction] temp: %v, humidity: %v, barometricPressure: %v, soilTempShallow: %v, soilTempDeep: %v, soilWaterContentShallow: %v, soilWaterContentDeep: %v, soilElectricalConductivity: %v", *temp, *humidity, *barometricPressure, *soilTempShallow, *soilTempDeep, *soilWaterContentShallow, *soilWaterContentDeep, *soilElectricalConductivity)

		// algorithm input sequence:
		// soilElectricalConductivity,
		// humidity
		// barometricPressure,
		// soilTempShallow,
		// soilTempDeep,
		// soilWaterContentShallow,
		// soilWaterContentDeep,
		// temp
		requestURL := fmt.Sprintf("http://localhost:8080/stations?soil_electrical_conductivity=%v&humidity=%v&barometric_pressure=%v&soil_temp_shallow=%v&soil_temp_deep=%v&soil_water_content_shallow=%v&soil_water_content_deep=%v&temp=%v", *soilElectricalConductivity, *humidity, *barometricPressure, *soilTempShallow, *soilTempDeep, *soilWaterContentShallow, *soilWaterContentDeep, *temp)
		req, err := http.NewRequest(http.MethodGet, requestURL, nil)
		if err != nil {
			c.InternalServerError(errors.New("cannot create http request"))
			return
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			c.InternalServerError(errors.New("algo service unavailable"))
			return
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			c.InternalServerError(errors.New("failed to read res body"))
			return
		}
		var result service.StationPredictionResult
		if err := json.Unmarshal(body, &result); err != nil {
			c.InternalServerError(errors.New("failed to unmarshal JSON"))
			return
		}
		c.OK(result)
	} else {
		c.BadRequest(errors.New("missing params"))
	}
}
