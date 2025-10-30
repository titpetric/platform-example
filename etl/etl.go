// This example omits the platform/module package, explicitly or autoloaded.
// It adds a manually set up ETL package explicitly in the start() function.
package main

import (
	"log"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/titpetric/platform"
	"github.com/titpetric/platform-example/etl/internal"
)

func main() {
	// Register common middleware.
	platform.Use(middleware.Logger)

	if err := start(); err != nil {
		log.Fatalf("exit error: %v", err)
	}
}

func start() error {
	platform.Register(internal.NewHandler())

	return platform.Start()
}
