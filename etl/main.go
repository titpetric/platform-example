package main

import (
	"log"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/titpetric/platform"
	"github.com/titpetric/platform/registry"

	"github.com/titpetric/platform-example/etl/internal"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "modernc.org/sqlite"
)

var _ registry.Module = (*internal.Handler)(nil)

func main() {
	if err := start(); err != nil {
		log.Fatalf("exit error: %v", err)
	}
}

func start() error {
	registry.AddMiddleware(middleware.Logger)

	etl, err := internal.NewHandler()
	if err != nil {
		return err
	}
	registry.AddModule(etl)

	if err := platform.Start(); err != nil {
		return err
	}
	return nil
}
