package app

import (
	"EnvMonitoringDashboard/api_src/config"
	"EnvMonitoringDashboard/api_src/logger"

	"github.com/gin-gonic/gin"
)

type App struct {
	*gin.Engine
	*config.Config
	*logger.Logger
}

func NewApp(
	e *gin.Engine,
	c *config.Config,
	l *logger.Logger,
) *App {
	return &App{e, c, l}
}
