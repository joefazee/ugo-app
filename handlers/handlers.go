package handlers

import (
	"net/http"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/joefazee/ladiwork/data"
	"github.com/joefazee/ugo"
)

type Handler struct {
	App    *ugo.Ugo
	Models data.Models
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	defer h.App.LoadTime(time.Now())
	err := h.App.Render.Page(w, r, "home", nil, nil)

	if err != nil {
		h.App.ErrorLog.Println("Error rendering home template", err)
	}
}