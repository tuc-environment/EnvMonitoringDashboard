package controller

import (
	"api/api_src/config"
	"api/api_src/logger"
	"api/api_src/service"

	"github.com/gin-gonic/gin"
)

type StationAPI struct {
	config         *config.Config
	logger         *logger.Logger
	accountService *service.StationService
}

func NewStationAPI(c *config.Config, l *logger.Logger, s *service.StationService) *StationAPI {
	return &StationAPI{c, l, s}
}

// Register godoc
//
//	@Summary		Get stations
//	@Description	Get stations with params
//	@Tags			stations
//	@Accept			json
//	@Produce		json
//	@Success		200	"Return station json array"
//	@Router			/stations [get]
func (api *StationAPI) GetStations(g *gin.Context) {
	c := WrapContext(g)
	stations, err := api.accountService.GetStations()
	if err != nil {
		c.BadRequest(err)
		return
	}
	c.OK(stations)
}
