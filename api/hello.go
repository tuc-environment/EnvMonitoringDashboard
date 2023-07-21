package api

import (
	"api/config"
	"api/db"

	"github.com/gin-gonic/gin"
)

type HelloAPI struct {
	c *config.Config
	s *db.Service
}

func NewHelloAPI(c *config.Config, s *db.Service) *HelloAPI {
	return &HelloAPI{c, s}
}

func (api *HelloAPI) Ping(g *gin.Context) {
	c := WrapContext(g)
	c.OK("pong")
}
