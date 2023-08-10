package main

import (
	"EnvMonitoringDashboard/api"
	"embed"
	"fmt"
	"io/fs"
)

//go:embed _webapp/dist
var embedWebappFS embed.FS

//	@title			天津商业大学环境监测系统API
//	@version		1.0
//	@description	天津商业大学环境监测系统API文档
//	@termsOfService	https://www.tjcu.edu.cn/

//	@host		localhost:8080
//	@BasePath	/

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	webappFS, err := fs.Sub(embedWebappFS, "_webapp/dist")
	if err != nil {
		panic(err)
	}

	app := api.NewApp(webappFS)
	log := app.Sugar()
	defer log.Sync()
	log.Infoln("Starting server on port", app.PORT)
	app.Run(fmt.Sprintf(":%d", app.PORT))
}
