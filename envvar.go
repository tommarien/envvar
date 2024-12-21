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
	"fmt"
	"os"
	"strconv"
)

// NotSetError represents an error when a required environment variable is not set.
type NotSetError struct {
	Func, Key string
}

func (e NotSetError) Error() string {
	return fmt.Sprintf("envvar.%s: missing required environment variable %q", e.Func, e.Key)
}

// ConversionError represents an error when a conversion of an environment variable fails.
type ConversionError struct {
	Func, Key, Value string
	err              error
}

func (e ConversionError) Error() string {
	return fmt.Sprintf("envvar.%s: invalid value %q for environment variable %q: %v", e.Func, e.Value, e.Key, e.err)
}

func (e ConversionError) Unwrap() error {
	return e.err
}

// String retrieves the value of the environment variable named by the key.
// If the variable is present in the environment the value (which may be empty) is returned
// otherwise it returns a [NotSetError].
func String(key string) (string, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return "", NotSetError{
			Func: "String",
			Key:  key,
		}
	}

	return value, nil
}

// StringOrDefault retrieves the value of the environment variable named by the key.
// If the variable is present in the environment the value (which may be empty) is returned
// otherwise it returns the defaultValue.
func StringOrDefault(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return value
}

// Int retrieves the value of the environment variable named by the key and converts it to an integer.
// If the variable is present in the environment the value is returned otherwise it returns a [NotSetError].
// If the conversion fails it returns a [ConversionError].
func Int(key string) (int, error) {
	variable, ok := os.LookupEnv(key)
	if !ok {
		return 0, NotSetError{Func: "Int", Key: key}
	}

	value, err := strconv.Atoi(variable)
	if err != nil {
		return 0, ConversionError{Func: "Int", Key: key, Value: variable, err: err}
	}

	return value, nil
}
