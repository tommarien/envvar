package envvar

import (
	"errors"
	"fmt"
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
	Err              error
}

func (e ConversionError) Error() string {
	return fmt.Sprintf("envvar.%s: invalid value %q for environment variable %q: %v", e.Func, e.Value, e.Key, e.Err)
}

func (e ConversionError) Unwrap() error {
	return e.Err
}

// ParserError represents an error when there are errors
// parsing the environment variables using [Parser].
type ParserError struct {
	errors map[string]error
}

// GetError returns the error for the given key if any.
func (p ParserError) GetError(key string) error {
	return p.errors[key]
}

func (p ParserError) Error() string {
	return fmt.Sprintf("failed to parse environment variables:\n%v", errors.Join(p.Unwrap()...))
}

func (p ParserError) Unwrap() []error {
	errs := make([]error, 0, len(p.errors))

	for _, err := range p.errors {
		errs = append(errs, err)
	}

	return errs
}
