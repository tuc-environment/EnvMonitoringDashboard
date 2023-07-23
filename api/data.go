package api

import "github.com/gin-gonic/gin"

type DataAPI struct {
}

func NewDataAPI() *DataAPI {
	return &DataAPI{}
}

func (api *DataAPI) Ping(g *gin.Context) {
	g.String(200, "pong")
}
