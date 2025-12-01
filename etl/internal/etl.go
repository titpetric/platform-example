package internal

import (
	"context"

	etl "github.com/titpetric/etl/server"
	"github.com/titpetric/platform"
)

type Handler struct {
	platform.UnimplementedModule
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Name() string {
	return "etl"
}

func (h *Handler) Mount(_ context.Context, r platform.Router) error {
	handler, err := etl.NewHandler()
	if err != nil {
		return err
	}

	r.Mount("/etl/", handler)

	return nil
}
