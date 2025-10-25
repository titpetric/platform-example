# ETL

The etl tool at
[github.com/titpetric/etl](https://github.com/titpetric/etl) provides a
server package, which exposes a HTTP handler. As a http handler can be
attached to a chi route, we can integrate the etl package with
[titpetric/platform](https://github.com/titpetric/platform).

There are two ways to integrate it:

- Directly by modifying your application entrypoint, [etl.go](./etl.go)
- Indirectly by using an autoloader package and `init`, [etl_autoload.go](./etl_autoload.go)

The convention picked up is based on
[joho/godotenv](https://github.com/joho/godotenv). Their package
provides an `autoload` function that declares the necessary `init` hook.

In both cases, platform/registry `Add` or `AddModule` will register
a new plugin into the platform. Autoloading is used to avoid needing
to modify the `main()` function further.