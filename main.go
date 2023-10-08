package main

import "gitlab.com/hbarral/regius"

type application struct {
	App *regius.Regius
}

func main() {
	r := initApplication()
	r.App.ListenAndServe()
}
