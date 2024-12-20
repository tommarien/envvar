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
)

// NotSetError represents an error when a required environment variable is not set.
type NotSetError struct {
	Key string
}

func (e NotSetError) Error() string {
	return fmt.Sprintf("missing required environment variable %q", e.Key)
}

// String retrieves the value of the environment variable named by the key.
// If the variable is present in the environment the value (which may be empty) is returned
// otherwise it returns a [NotSetError].
func String(key string) (string, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return "", NotSetError{Key: key}
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
