package api

import (
	"api/api_src"
	"api/api_src/app"
	"io/fs"
	"net/http"
)

func NewApp(webappFS fs.FS) *app.App {
	app, err := api_src.InitializeApp()
	if err != nil {
		panic(err)
	}
	return app
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app := NewApp(nil)
	app.ServeHTTP(w, r)
}
