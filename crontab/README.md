# Crontab

In [internal/crontab.go](./internal/crontab.go), you'll find a module that implements
the following bits:

## Types

```go
type Crontab struct {
	scheduler *cron.Cron
}
```

## Function symbols

- `func NewCrontab () (*Crontab, error)`
- `func (*Crontab) Close ()`
- `func (*Crontab) Mount ( chi.Router)`
- `func (*Crontab) Name () string`

# Example usage

This is how a module can be used explicitly, as shown in [main.go](./main.go);

```go
func start(ctx context.Context) error {
	// Create the module instance.
	crontab, err := internal.NewCrontab()
	if err != nil {
		return err
	}

	// Create a platform instance.
	svc, err := platform.New()
	if err != nil {
		return err
	}

	// Add common middleware.
	svc.AddMiddleware(middleware.Logger)

	// Add the crontab module. Other modules could be added.
	svc.AddModule(crontab)

	// Start the server and wait for an exit.
	svc.Serve(ctx)
	svc.Wait()

	return nil
}
```

The platform allows you to use middleare in two locations, for all your
modules, or within the `Mount()` function of the module to apply only to
the routes of your module.

The crontab example is a good case demonstrating lifecycle control. A
Platform object will invoke the `Mount` function for the module from the
platforms `Serve` function.