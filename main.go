package main

import (
	"regius-app/data"
	"regius-app/handlers"
	"regius-app/middleware"

	"gitlab.com/hbarral/regius"
)

type application struct {
	App        *regius.Regius
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	r := initApplication()
	r.App.ListenAndServe()
}
