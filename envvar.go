// Package envvar reads and converts environment variables, with no
// dependencies outside the standard library.
//
// Every accessor takes the name of an environment variable and reports a
// [NotSetError] when it is not set, or a [ParseError] when its value cannot
// be converted:
//
//	port, err := envvar.Int("PORT")
//
// Each has an OrDefault counterpart falling back to the given value when the
// variable is not set. [StringOrDefault] is the exception that cannot fail,
// so it returns no error:
//
//	region := envvar.StringOrDefault("REGION", "eu-west-1")
//	timeout, err := envvar.DurationOrDefault("TIMEOUT", 10*time.Minute)
//
// [String], [Int], [Bool], [UInt], [Float] and [Duration] cover the common
// types. [Func] and [FuncOrDefault] take a [ParseFunc] and so convert into
// any type.
//
// # Parsing several variables at once
//
// A [VarSet] collects variables and parses them together, so a single call
// reports every problem instead of only the first. Its zero value is ready
// for use:
//
//	var (
//		cfg Config
//		set envvar.VarSet
//	)
//
//	envvar.StringVar(&set, &cfg.Name, "NAME")
//	envvar.StringVarOrDefault(&set, &cfg.Region, "REGION", "eu-west-1")
//	envvar.IntVar(&set, &cfg.Port, "PORT")
//	envvar.DurationVarOrDefault(&set, &cfg.Timeout, "TIMEOUT", 10*time.Minute)
//
//	if err := set.Parse(); err != nil {
//		log.Fatal(err)
//	}
//
// [VarSet.Parse] runs the registrations in order, assigns each value in place
// and returns a [VarSetParseError] holding every failure. Each accessor has a
// matching registration function, [FuncVar] and [FuncVarOrDefault] included.
package envvar

import (
	"errors"
	"os"
	"strconv"
	"time"
)

// ParseFunc converts a string value into T.
type ParseFunc[T any] func(value string) (T, error)

// Func returns the value of the environment variable named by key, converted by parse.
// If it is not set, it returns a [NotSetError].
// If parse fails, it returns a [ParseError].
func Func[T any](key, funcName string, parse ParseFunc[T]) (T, error) {
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

// FuncOrDefault returns the value of the environment variable named by key, converted by parse.
// If it is not set, it returns defaultValue instead.
// If parse fails, it returns a [ParseError].
func FuncOrDefault[T any](key string, defaultValue T, funcName string, parse ParseFunc[T]) (T, error) {
	value, err := Func(key, funcName, parse)
	var notSetErr NotSetError
	if errors.As(err, &notSetErr) {
		return defaultValue, nil
	}

	return value, err
}

// String returns the value of the environment variable named by key.
// If it is not set, it returns a [NotSetError].
func String(key string) (string, error) {
	return Func(key, "String", func(v string) (string, error) { return v, nil })
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
	return Func(key, "Int", strconv.Atoi)
}

// IntOrDefault returns the value of the environment variable named by key, converted to an integer.
// If it is not set, it returns defaultValue instead.
// If the conversion fails, it returns a [ParseError].
func IntOrDefault(key string, defaultValue int) (int, error) {
	return FuncOrDefault(key, defaultValue, "IntOrDefault", strconv.Atoi)
}

// Bool returns the value of the environment variable named by key, converted to a boolean.
// If it is not set, it returns a [NotSetError].
// If the conversion fails, it returns a [ParseError].
func Bool(key string) (bool, error) {
	return Func(key, "Bool", strconv.ParseBool)
}

// BoolOrDefault returns the value of the environment variable named by key, converted to a boolean.
// If it is not set, it returns defaultValue instead.
// If the conversion fails, it returns a [ParseError].
func BoolOrDefault(key string, defaultValue bool) (bool, error) {
	return FuncOrDefault(key, defaultValue, "BoolOrDefault", strconv.ParseBool)
}

func parseUint(value string) (uint, error) {
	v, err := strconv.ParseUint(value, 10, 0)
	return uint(v), err
}

// UInt returns the value of the environment variable named by key, converted to an unsigned integer.
// If it is not set, it returns a [NotSetError].
// If the conversion fails, it returns a [ParseError].
func UInt(key string) (uint, error) {
	return Func(key, "UInt", parseUint)
}

// UIntOrDefault returns the value of the environment variable named by key, converted to an unsigned integer.
// If it is not set, it returns defaultValue instead.
// If the conversion fails, it returns a [ParseError].
func UIntOrDefault(key string, defaultValue uint) (uint, error) {
	return FuncOrDefault(key, defaultValue, "UIntOrDefault", parseUint)
}

func parseFloat(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}

// Float returns the value of the environment variable named by key, converted to a float64.
// If it is not set, it returns a [NotSetError].
// If the conversion fails, it returns a [ParseError].
func Float(key string) (float64, error) {
	return Func(key, "Float", parseFloat)
}

// FloatOrDefault returns the value of the environment variable named by key, converted to a float64.
// If it is not set, it returns defaultValue instead.
// If the conversion fails, it returns a [ParseError].
func FloatOrDefault(key string, defaultValue float64) (float64, error) {
	return FuncOrDefault(key, defaultValue, "FloatOrDefault", parseFloat)
}

// Duration returns the value of the environment variable named by key, converted to a [time.Duration].
// If it is not set, it returns a [NotSetError].
// If the conversion fails, it returns a [ParseError].
func Duration(key string) (time.Duration, error) {
	return Func(key, "Duration", time.ParseDuration)
}

// DurationOrDefault returns the value of the environment variable named by key, converted to a [time.Duration].
// If it is not set, it returns defaultValue instead.
// If the conversion fails, it returns a [ParseError].
func DurationOrDefault(key string, defaultValue time.Duration) (time.Duration, error) {
	return FuncOrDefault(key, defaultValue, "DurationOrDefault", time.ParseDuration)
}
