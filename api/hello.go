package api

import (
	"api/config"
	"api/service"

	"github.com/gin-gonic/gin"
)

type HelloAPI struct {
	config         *config.Config
	accountService *service.AccountService
}

func NewHelloAPI(c *config.Config, s *service.AccountService) *HelloAPI {
	return &HelloAPI{c, s}
}

func (api *HelloAPI) Ping(g *gin.Context) {
	c := WrapContext(g)
	c.OK("pong")
}
