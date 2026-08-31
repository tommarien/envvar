# envvar
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/tommarien/envvar/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/tommarien/envvar/tree/main)

A zero dependency library to parse environment variables

## Install

```sh
go get github.com/tommarien/envvar
```

## Usage

Read a single environment variable with one of the accessors. Each returns a
`NotSetError` when the variable is not set, and a `ParseError` when its value
cannot be converted.

```go
port, err := envvar.Int("PORT")
```

Every accessor has an `OrDefault` counterpart that falls back to the given
value when the variable is not set. Only `StringOrDefault` cannot fail, so it
returns no error.

```go
region := envvar.StringOrDefault("REGION", "eu-west-1")
timeout, err := envvar.DurationOrDefault("TIMEOUT", 10*time.Minute)
```

| Type            | Required          | With default               |
| --------------- | ----------------- | -------------------------- |
| `string`        | `String`          | `StringOrDefault`          |
| `int`           | `Int`             | `IntOrDefault`             |
| `bool`          | `Bool`            | `BoolOrDefault`            |
| `uint`          | `UInt`            | `UIntOrDefault`            |
| `float64`       | `Float`           | `FloatOrDefault`           |
| `time.Duration` | `Duration`        | `DurationOrDefault`        |
| any other type  | `Func`            | `FuncOrDefault`            |

## Parsing several variables at once

A `VarSet` collects variables and parses them together, so one call reports
every problem instead of only the first. Its zero value is ready for use.

```go
type Config struct {
	Name    string
	Region  string
	Port    int
	Timeout time.Duration
}

var (
	cfg Config
	set envvar.VarSet
)

envvar.StringVar(&set, &cfg.Name, "NAME")
envvar.StringVarOrDefault(&set, &cfg.Region, "REGION", "eu-west-1")
envvar.IntVar(&set, &cfg.Port, "PORT")
envvar.DurationVarOrDefault(&set, &cfg.Timeout, "TIMEOUT", 10*time.Minute)

if err := set.Parse(); err != nil {
	log.Fatal(err)
}
```

`Parse` runs the registrations in order and assigns each value in place. If
any of them fail it returns a `VarSetParseError` holding every failure, which
you can inspect per variable:

```go
var parseErr envvar.VarSetParseError
if errors.As(err, &parseErr) {
	fmt.Println(parseErr.GetError("PORT"))
}
```

Each accessor has a matching registration function: `StringVar`, `IntVar`,
`BoolVar`, `UIntVar`, `FloatVar`, `DurationVar` and their `OrDefault`
counterparts.

## Custom types

`Func` and `FuncOrDefault` take a `ParseFunc`, which converts the raw value
into any type you need.

```go
type Level int

const (
	LevelInfo Level = iota + 1
	LevelDebug
)

func parseLevel(value string) (Level, error) {
	switch value {
	case "info":
		return LevelInfo, nil
	case "debug":
		return LevelDebug, nil
	default:
		return 0, errors.New("unknown level")
	}
}

level, err := envvar.FuncOrDefault("LOG_LEVEL", LevelInfo, "LogLevel", parseLevel)
```

The same `ParseFunc` registers with a `VarSet` through `FuncVar` and
`FuncVarOrDefault`.

```go
envvar.FuncVarOrDefault(&set, &cfg.Level, "LOG_LEVEL", LevelInfo, "LogLevel", parseLevel)
```

The `funcName` argument is what the returned `NotSetError` and `ParseError`
report as the failing function.
