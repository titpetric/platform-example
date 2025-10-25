// This example omits the platform/module package, explicitly or autoloaded.
// It adds a manually set up ETL package explicitly in the start() function.
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

func main() {
	// Register common middleware.
	registry.AddMiddleware(middleware.Logger)

	if err := start(); err != nil {
		log.Fatalf("exit error: %v", err)
	}
}

func start() error {
	etl, err := internal.NewHandler()
	if err != nil {
		return err
	}
	registry.AddModule(etl)

	return platform.Start()
}
