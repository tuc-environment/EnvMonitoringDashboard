package app

import (
	"api/config"
	"api/controller"
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
	accountAPI *controller.AccountAPI,
	stationAPI *controller.StationAPI,
	recordAPI *controller.RecordAPI,
	dataAPI *controller.DataAPI,
) (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()
	e.Use(NewLoggerMiddleware(c, log), gin.Recovery())

	e.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Authorization", "Origin", "User-Agent", "X-API-Key"},
		ExposeHeaders:    []string{"Content-Type", "Content-Length", "Authorization", "Origin", "User-Agent", "X-API-Key"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := e.Group("/api")

	// public apis
	api.GET("/", dataAPI.Ping)
	api.POST("/register", accountAPI.Register)
	api.POST("/login", accountAPI.Login)
	api.GET("/stations", stationAPI.GetStations)

	// authorised required apis
	api.Use(accountAPI.AuthMiddleware)
	{
		api.GET("/account", accountAPI.GetAccount)
	}

	// append docs
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return e, nil
}
