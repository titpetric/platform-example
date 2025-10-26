package autoload

import (
	"log"

	"github.com/titpetric/platform"
	"github.com/titpetric/platform-example/etl/internal"
)

func init() {
	m, err := internal.NewHandler()
	if err != nil {
		log.Fatalf("init error loading %s: %v", m.Name(), err)
	}
	platform.AddModule(m)
}
