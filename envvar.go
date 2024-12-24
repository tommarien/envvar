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
// If the conversion fails it returns a [ParseError].
func Int(key string) (int, error) {
	variable, ok := os.LookupEnv(key)
	if !ok {
		return 0, NotSetError{Func: "Int", Key: key}
	}

	value, err := strconv.Atoi(variable)
	if err != nil {
		return 0, ParseError{Func: "Int", Key: key, Value: variable, Err: err}
	}

	return value, nil
}

// IntOrDefault retrieves the value of the environment variable named by the key and converts it to an integer.
// If the variable is present in the environment the value is converted otherwise it returns the defaultValue.
// If the conversion fails it returns a [ParseError].
func IntOrDefault(key string, defaultValue int) (int, error) {
	variable, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue, nil
	}

	value, err := strconv.Atoi(variable)
	if err != nil {
		return 0, ParseError{Func: "IntOrDefault", Key: key, Value: variable, Err: err}
	}

	return value, nil
}
