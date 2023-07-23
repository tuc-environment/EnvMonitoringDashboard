package app

import (
	"api/config"
	"api/logger"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func NewLoggerMiddleware(c *config.Config, logger *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.Sugar()
		defer log.Sync()
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		latencyMsg := fmt.Sprintf("%8dms", latency.Milliseconds())
		log.Infoln(c.Request.Method, c.Writer.Status(), latencyMsg, c.Request.URL)
	}
}
