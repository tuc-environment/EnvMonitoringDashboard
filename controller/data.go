package controller

import "github.com/gin-gonic/gin"

type DataAPI struct {
}

func NewDataAPI() *DataAPI {
	return &DataAPI{}
}

// Basic Data godoc
//
//	@Summary		Get data
//	@Description	Get data
//	@Tags			api
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200	"Return data"
//	@Failure		401	"Token is incorrect"
//	@Router			/api [get]
func (api *DataAPI) Ping(g *gin.Context) {
	c := WrapContext(g)

	c.OK("pong")
}
