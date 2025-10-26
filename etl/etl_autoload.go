// This example autoloads the platform/module package.
// It also adds autoloading of the ETL package.
// The resulting main() function is identical to the upstream project.
package main

import (
	"log"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/titpetric/platform"
	"github.com/titpetric/platform/registry"

	// ETL requires sqlite driver to be loaded.
	_ "modernc.org/sqlite"

	// Registers the default platform modules.
	_ "github.com/titpetric/platform/module/autoload"

	// Registers the etl module.
	_ "github.com/titpetric/platform-example/etl/internal/autoload"
)

func main() {
	// Register common middleware.
	registry.AddMiddleware(middleware.Logger)

	if err := platform.Start(); err != nil {
		log.Fatalf("exit error: %v", err)
	}
}
