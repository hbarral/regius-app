package main

import (
	"log"
	"os"
	"regius-app/handlers"

	"gitlab.com/hbarral/regius"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	reg := &regius.Regius{}
	err = reg.New(path)

	if err != nil {
		log.Fatal(err)
	}

	reg.AppName = "regius-app"

	myHandlers := &handlers.Handlers{
		App: reg,
	}

	app := &application{
		App:      reg,
		Handlers: myHandlers,
	}

	app.App.Routes = app.routes()

	return app
}
