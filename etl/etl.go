// This example omits the platform/pkg/drivers package.
// It adds a manually set up ETL package explicitly in the start() function.
package main

import (
	"context"
	"log"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/titpetric/platform"
	"github.com/titpetric/platform-example/etl/internal"
)

func main() {
	if err := start(context.Background()); err != nil {
		log.Fatalf("exit error: %v", err)
	}
}

func start(ctx context.Context) error {
	p := platform.New(nil)
	p.Use(middleware.Logger)
	p.Register(internal.NewHandler())

	err := p.Start(ctx)
	if err != nil {
		return err
	}

	p.Wait()
	return nil
}
