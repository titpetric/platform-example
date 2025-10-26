package main

import (
	"context"
	"log"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/titpetric/platform"
	"github.com/titpetric/platform-example/crontab/internal"
)

func main() {
	ctx := context.Background()
	if err := start(ctx); err != nil {
		log.Fatalf("exit error: %v", err)
	}
}

func start(ctx context.Context) error {
	crontab, err := internal.NewCrontab()
	if err != nil {
		return err
	}

	svc, err := platform.New()
	if err != nil {
		return err
	}

	svc.AddMiddleware(middleware.Logger)
	svc.AddModule(crontab)

	svc.Serve(ctx)
	svc.Wait()

	return nil
}
