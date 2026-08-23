package envvar

import (
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
	envVar, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue, nil
	}

	value, err := parse(envVar)
	if err != nil {
		var zero T
		return zero, ParseError{Func: funcName, Key: key, Value: envVar, Err: err}
	}

	return value, nil
}

// String retrieves the value of the environment variable named by the key.
// If the variable is present in the environment, the value (which may be empty) is returned.
// If the variable is not present, it returns a [NotSetError].
func String(key string) (string, error) {
	return get(key, "String", func(v string) (string, error) { return v, nil })
}

// StringOrDefault retrieves the value of the environment variable named by the key.
// If the variable is present in the environment, the value (which may be empty) is returned.
// If the variable is not present, defaultValue is returned instead.
func StringOrDefault(key, defaultValue string) string {
	envVar, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return envVar
}

// Int retrieves the value of the environment variable named by the key and converts it to an integer.
// If the variable is present in the environment, the value is returned.
// If the variable is not present, it returns a [NotSetError].
// If the conversion fails it returns a [ParseError].
func Int(key string) (int, error) {
	return get(key, "Int", strconv.Atoi)
}

// IntOrDefault retrieves the value of the environment variable named by the key and converts it to an integer.
// If the variable is present in the environment, the value is returned.
// If the variable is not present, defaultValue is returned instead.
// If the conversion fails it returns a [ParseError].
func IntOrDefault(key string, defaultValue int) (int, error) {
	return getOrDefault(key, defaultValue, "IntOrDefault", strconv.Atoi)
}
