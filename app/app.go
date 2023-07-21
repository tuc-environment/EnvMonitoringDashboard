package app

import (
	"api/config"

	"github.com/gin-gonic/gin"
)

type App struct {
	*gin.Engine
	*config.Config
}

func NewApp(
	e *gin.Engine,
	c *config.Config,
) *App {
	return &App{e, c}
}
