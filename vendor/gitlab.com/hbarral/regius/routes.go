package regius

import (
	"fmt"
	"net/http"

	chi "github.com/go-chi/chi/v5"

	middleware "github.com/go-chi/chi/v5/middleware"
)

func (r *Regius) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)

	if r.Debug {
		mux.Use(middleware.Logger)

	}

	mux.Use(middleware.Recoverer)

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to Regius")
	})

	return mux
}
