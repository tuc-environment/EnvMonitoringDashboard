package app

import (
	"api/config"
	"api/logger"

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
