package envvar

import (
	"fmt"
)

// NotSetError represents an error when a required environment variable is not set.
type NotSetError struct {
	Func, Key string
}

func (e NotSetError) Error() string {
	return fmt.Sprintf("envvar.%s: missing required environment variable %q", e.Func, e.Key)
}

// ParseError represents an error when a conversion of an environment variable fails.
type ParseError struct {
	Func, Key, Value string
	Err              error
}

func (e ParseError) Error() string {
	return fmt.Sprintf("envvar.%s: invalid value %q for environment variable %q: %v", e.Func, e.Value, e.Key, e.Err)
}

func (e ParseError) Unwrap() error {
	return e.Err
}
