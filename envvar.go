package envvar

import (
	"errors"
	"os"
	"strconv"
)

// ParseFunc converts a raw environment variable value into T.
type ParseFunc[T any] func(value string) (T, error)

func get[T any](key, funcName string, parse ParseFunc[T]) (T, error) {
	var zero T
	envVar, ok := os.LookupEnv(key)
	if !ok {
		return zero, NotSetError{Func: funcName, Key: key}
	}

	value, err := parse(envVar)
	if err != nil {
		return zero, ParseError{Func: funcName, Key: key, Value: envVar, Err: err}
	}

	return value, nil
}

func getOrDefault[T any](key string, defaultValue T, funcName string, parse ParseFunc[T]) (T, error) {
	value, err := get(key, funcName, parse)
	if _, ok := errors.AsType[NotSetError](err); ok {
		return defaultValue, nil
	}

	return value, err
}

// String returns the value of the environment variable named by key.
// If it is not set, it returns a [NotSetError].
func String(key string) (string, error) {
	return get(key, "String", func(v string) (string, error) { return v, nil })
}

// StringOrDefault returns the value of the environment variable named by key.
// If it is not set, it returns defaultValue instead.
func StringOrDefault(key, defaultValue string) string {
	envVar, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return envVar
}

// Int returns the value of the environment variable named by key, converted to an integer.
// If it is not set, it returns a [NotSetError].
// If the conversion fails, it returns a [ParseError].
func Int(key string) (int, error) {
	return get(key, "Int", strconv.Atoi)
}

// IntOrDefault returns the value of the environment variable named by key, converted to an integer.
// If it is not set, it returns defaultValue instead.
// If the conversion fails, it returns a [ParseError].
func IntOrDefault(key string, defaultValue int) (int, error) {
	return getOrDefault(key, defaultValue, "IntOrDefault", strconv.Atoi)
}
