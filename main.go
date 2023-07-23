package main

import "fmt"

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
