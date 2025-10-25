package internal

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	etl "github.com/titpetric/etl/server"
)

type Handler struct {
	handler http.Handler
}

func NewHandler() (*Handler, error) {
	handler, err := etl.NewHandler()
	if err != nil {
		return nil, err
	}
	return &Handler{handler}, nil
}

func (h *Handler) Name() string {
	return "etl"
}

func (h *Handler) Mount(r chi.Router) {
	r.Mount("/etl/", h.handler)
}

func (*Handler) Close() {
}
