package main

import (
	"log"
	"os"

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
	reg.Debug = true

	app := &application{
		App: reg,
	}
	return app
}
