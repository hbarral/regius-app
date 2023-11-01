package handlers

import (
	"net/http"

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
