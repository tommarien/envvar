// Package envvar is a simple, zero-dependencies library to parse environment
// variables.
//
// Example:
//
//	type config struct {
//		Home string `env:"HOME"`
//	}
//	// parse
//	var cfg config
//	err := env.Parse(&cfg)
//	// or parse with generics
//	cfg, err := env.ParseAs[config]()
//
// Check the examples and README for more detailed usage.
package envvar

import (
	"os"
	"strconv"
)

type config[T any] struct {
	name         string
	defaultValue *T
}

func withName[T any](name string) func(c *config[T]) {
	return func(c *config[T]) {
		c.name = name
	}
}

// WithDefault sets the default value for the environment variable
// if it is not set.
func WithDefault[T any](defaultValue T) func(c *config[T]) {
	return func(c *config[T]) {
		c.defaultValue = &defaultValue
	}
}

func zero[T any]() T {
	var zero T
	return zero
}

func get[T any](key string, parser func(string) (T, error), options ...func(c *config[T])) (T, error) {
	cfg := config[T]{
		name: "Get",
	}

	for _, opt := range options {
		opt(&cfg)
	}

	envVar, ok := os.LookupEnv(key)
	if !ok {
		if cfg.defaultValue != nil {
			return *cfg.defaultValue, nil
		}

		return zero[T](), NotSetError{Func: cfg.name, Key: key}
	}

	value, err := parser(envVar)
	if err != nil {
		return zero[T](), ParseError{Func: cfg.name, Key: key, Value: envVar, Err: err}
	}

	return value, nil
}

// String retrieves the value of the environment variable named by the key.
// If the variable is present in the environment, the value (which may be empty) is returned.
// If the variable is not present, it returns a [NotSetError] or the default value if it was set via options.
func String(key string, options ...func(c *config[string])) (string, error) {
	return get(
		key,
		func(s string) (string, error) { return s, nil },
		append(options, withName[string]("String"))...,
	)
}

// Int retrieves the value of the environment variable named by the key and converts it to an integer.
// If the variable is present in the environment, the value is returned.
// If the variable is not present, it returns a [NotSetError] or the default value if it was set via options.
// If the conversion fails it returns a [ParseError].
func Int(key string, options ...func(c *config[int])) (int, error) {
	return get(
		key,
		strconv.Atoi,
		append(options, withName[int]("Int"))...,
	)
}
