package handlers

import (
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"gitlab.com/hbarral/regius"
)

type Handlers struct {
	App *regius.Regius
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "home", nil, nil)

	if err != nil {
		h.App.ErrorLog.Println("error rendering", err)
	}
}

func (h *Handlers) GoPage(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.GoPage(w, r, "home", nil)

	if err != nil {
		h.App.ErrorLog.Println("error rendering", err)
	}
}

func (h *Handlers) JetPage(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "jet-template", nil, nil)

	if err != nil {
		h.App.ErrorLog.Println("error rendering", err)
	}
}

func (h *Handlers) SessionTest(w http.ResponseWriter, r *http.Request) {
	data := "bar"

	h.App.Session.Put(r.Context(), "foo", data)

	value := h.App.Session.GetString(r.Context(), "foo")

	vars := make(jet.VarMap)
	vars.Set("foo", value)

	err := h.App.Render.JetPage(w, r, "sessions", vars, nil)

	if err != nil {
		h.App.ErrorLog.Println("error rendering", err)
	}
}
