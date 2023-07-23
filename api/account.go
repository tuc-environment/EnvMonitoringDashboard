package api

import (
	"api/config"
	"api/logger"
	"api/service"

	"github.com/gin-gonic/gin"
)

type AccountAPI struct {
	config         *config.Config
	logger         *logger.Logger
	accountService *service.AccountService
}

func NewAccountAPI(c *config.Config, l *logger.Logger, s *service.AccountService) *AccountAPI {
	return &AccountAPI{c, l, s}
}

func (api *AccountAPI) Ping(g *gin.Context) {
	c := WrapContext(g)
	c.OK("pong")
}
