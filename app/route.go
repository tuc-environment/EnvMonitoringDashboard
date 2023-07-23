package app

import (
	"api/api"
	"api/config"
	"api/logger"
	"time"

	"github.com/gin-contrib/cors"

	_ "api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func NewEngine(
	c *config.Config,
	log *logger.Logger,
	accountAPI *api.AccountAPI,
) (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	e := gin.Default()
	e.Use(NewLoggerMiddleware(log), gin.Recovery())

	e.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Authorization", "Origin", "User-Agent", "X-API-Key"},
		ExposeHeaders:    []string{"Content-Type", "Content-Length", "Authorization", "Origin", "User-Agent", "X-API-Key"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// append docs
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	e.GET("/", accountAPI.Ping)

	return e, nil
}
