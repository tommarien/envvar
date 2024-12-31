// Package envvar is a simple, zero-dependencies library to parse environment
// variables.
//
// Example:
//
//	type Config struct {
//		Port  int
//	}
//
//	var config Config
//
//	group := envvar.NewGroup()
//	group.IntVar(&config.Port, "PORT", envvar.WithDefault(8080))
//	err := group.Parse()
//
//	if err != nil {
//	  log.Fatal(err)
//	}
//
// Check the examples and README for more detailed usage.
package envvar

import (
	"os"
	"strconv"
)

type parseConfig[T any] struct {
	name         string
	defaultValue *T
}

func withName[T any](name string) func(c *parseConfig[T]) {
	return func(c *parseConfig[T]) {
		c.name = name
	}
}

// WithDefault sets the default value for the environment variable
// if it is not set.
func WithDefault[T any](defaultValue T) func(c *parseConfig[T]) {
	return func(c *parseConfig[T]) {
		c.defaultValue = &defaultValue
	}
}

func zero[T any]() T {
	var zero T
	return zero
}

// TODO: Rename to Func
func get[T any](key string, parser func(string) (T, error), options ...func(c *parseConfig[T])) (T, error) {
	cfg := parseConfig[T]{
		name: "Func",
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
func String(key string, options ...func(c *parseConfig[string])) (string, error) {
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
func Int(key string, options ...func(c *parseConfig[int])) (int, error) {
	return get(
		key,
		strconv.Atoi,
		append(options, withName[int]("Int"))...,
	)
}

// Bool retrieves the value of the environment variable named by the key and converts it to a boolean.
// If the variable is present in the environment, the value is returned.
// If the variable is not present, it returns a [NotSetError] or the default value if it was set via options.
// If the conversion fails it returns a [ParseError].
func Bool(key string, options ...func(c *parseConfig[bool])) (bool, error) {
	return get(
		key,
		strconv.ParseBool,
		append(options, withName[bool]("Bool"))...,
	)
}
