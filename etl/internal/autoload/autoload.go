package autoload

import (
	"github.com/titpetric/platform"
	"github.com/titpetric/platform-example/etl/internal"
)

func init() {
	platform.Register(internal.NewHandler())
}
