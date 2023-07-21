package api

import (
	"api/config"

	"github.com/gin-gonic/gin"
)

type HelloAPI struct {
	c *config.Config
}

func NewHelloAPI(c *config.Config) *HelloAPI {
	return &HelloAPI{c}
}

func (api *HelloAPI) Ping(g *gin.Context) {
	c := WrapContext(g)
	c.OK("pong")
}
