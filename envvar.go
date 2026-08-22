package envvar

import (
	"os"
	"strconv"
)

// String retrieves the value of the environment variable named by the key.
// If the variable is present in the environment, the value (which may be empty) is returned.
// If the variable is not present, it returns a [NotSetError].
func String(key string) (string, error) {
	envVar, ok := os.LookupEnv(key)
	if !ok {
		return "", NotSetError{Func: "String", Key: key}
	}

	return envVar, nil
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
	envVar, ok := os.LookupEnv(key)
	if !ok {
		return 0, NotSetError{Func: "Int", Key: key}
	}

	value, err := strconv.Atoi(envVar)
	if err != nil {
		return 0, ParseError{Func: "Int", Key: key, Value: envVar, Err: err}
	}

	return value, nil
}

// IntOrDefault retrieves the value of the environment variable named by the key and converts it to an integer.
// If the variable is present in the environment, the value is returned.
// If the variable is not present, defaultValue is returned instead.
// If the conversion fails it returns a [ParseError].
func IntOrDefault(key string, defaultValue int) (int, error) {
	envVar, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue, nil
	}

	value, err := strconv.Atoi(envVar)
	if err != nil {
		return 0, ParseError{Func: "IntOrDefault", Key: key, Value: envVar, Err: err}
	}

	return value, nil
}
