# ETL

The etl tool at
[github.com/titpetric/etl](https://github.com/titpetric/etl) provides a
server package, which exposes a HTTP handler. As a http handler can be
attached to a chi route, we can integrate the etl package with
[titpetric/platform](https://github.com/titpetric/platform).

Meaningful file examples for explicit extension:

- [etl.go](./etl.go): add modules from logic in main
- [internal/etl.go](./internal/etl.go): the implementation of a module

And the example for automatic extension (dependency inversion).

- [etl_autoload.go](./etl_autoload.go): uses autoloading to register modules
- [internal/autoload/autoload.go](./internal/autoload/autoload.go): implement module autoloading

The convention picked up for automatic extension is based on
[joho/godotenv](https://github.com/joho/godotenv). The package
provides an `autoload` package that declares the necessary `init` hook.

In both cases, platform/registry `Add` or `AddModule` will register
a new plugin into the platform. Autoloading is used to avoid needing
to modify the `main()` function further. A single autoloading package
could load multiple modules, preventing further changes in `main.go`.