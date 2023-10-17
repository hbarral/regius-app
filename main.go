package main

import (
	"regius-app/handlers"

	"gitlab.com/hbarral/regius"
)

type application struct {
	App      *regius.Regius
	Handlers *handlers.Handlers
}

func main() {
	r := initApplication()
	r.App.ListenAndServe()
}
