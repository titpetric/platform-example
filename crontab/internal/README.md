# Package internal

```go
import (
	"github.com/titpetric/platform-example/crontab/internal"
}
```

## Types

```go
type Crontab struct {
	scheduler *cron.Cron
}
```

## Function symbols

- `func NewCrontab () (*Crontab, error)`
- `func (*Crontab) Close ()`
- `func (*Crontab) Mount (chi.Router)`
- `func (*Crontab) Name () string`

### NewCrontab

```go
func NewCrontab () (*Crontab, error)
```

### Close

```go
func (*Crontab) Close ()
```

### Mount

```go
func (*Crontab) Mount (chi.Router)
```

### Name

```go
func (*Crontab) Name () string
```


