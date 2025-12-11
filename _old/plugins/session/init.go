package session

import (
	"github.com/go-chi/chi/v5"

	v1 "github.com/titpetric/platform/proto/v1"
	"github.com/titpetric/platform/registry"
)

func init() {
	registry.Add("session", func(r chi.Router) {
		srv := v1.NewSessionServiceServer(NewService())

		r.Post(v1.SessionServicePathPrefix+"*", srv.ServeHTTP)
	}, nil)
}
