package session

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	v1 "github.com/titpetric/platform/proto/v1/v1connect"
	"github.com/titpetric/platform/registry"
)

func init() {
	registry.Add("session.client", func(r chi.Router) {
		routePath, handler := v1.NewClientServiceHandler(NewService())

		r.Post("/connect"+routePath+"*", http.StripPrefix("/connect/", handler).ServeHTTP)
	}, nil)
}
