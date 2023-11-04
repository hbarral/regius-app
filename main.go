package main

import (
	"regius-app/data"
	"regius-app/handlers"

	"gitlab.com/hbarral/regius"
)

type application struct {
	App      *regius.Regius
	Handlers *handlers.Handlers
	Models   data.Models
}

func main() {
	r := initApplication()
	r.App.ListenAndServe()
}
