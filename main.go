package main

import (
	"fmt"
)

//	@title			天津商业大学环境监测系统API
//	@version		1.0
//	@description	天津商业大学环境监测系统API文档
//	@termsOfService	https://www.tjcu.edu.cn/

//	@host		www.tjcu.edu.cn
//	@BasePath	/

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	app, err := InitializeApp()
	if err != nil {
		panic(err)
	}

	log := app.Sugar()
	defer log.Sync()
	log.Infoln("Starting server on port", app.PORT)
	app.Run(fmt.Sprintf(":%d", app.PORT))
}
