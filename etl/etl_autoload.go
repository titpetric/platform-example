// This example autoloads the platform/module package.
// It also adds autoloading of the ETL package.
// The resulting main() function is identical to the upstream project.
package main

import (
	"log"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/titpetric/platform"

	// Registers the default platform modules.
	_ "github.com/titpetric/platform/pkg/drivers"

	// Registers the etl module.
	_ "github.com/titpetric/platform-example/etl/internal/autoload"
)

func main() {
	// Register common middleware.
	platform.Use(middleware.Logger)

	p, err := platform.Start()
	if err != nil {
		log.Fatalf("exit error: %v", err)
	}
	p.Wait()
}
