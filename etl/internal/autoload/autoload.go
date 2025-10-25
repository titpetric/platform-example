package autoload

import (
	"log"

	"github.com/titpetric/platform-example/etl/internal"
	"github.com/titpetric/platform/registry"
)

func init() {
	m, err := internal.NewHandler()
	if err != nil {
		log.Fatalf("init error loading %s: %v", m.Name(), err)
	}
	registry.AddModule(m)
}
