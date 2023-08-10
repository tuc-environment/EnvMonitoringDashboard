package app

import (
	"api/api_src/config"
	"api/api_src/controller"
	"api/api_src/logger"
	"io/fs"
	"time"

	"github.com/gin-contrib/cors"

	_ "api/api_src/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func NewEngine(
	webappFS fs.FS,

	c *config.Config,
	log *logger.Logger,
	noRouteAPI *controller.NoRouteAPI,
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

	e.NoRoute(noRouteAPI.ServeWebapp(webappFS))

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
