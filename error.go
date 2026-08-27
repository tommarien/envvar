package envvar

import (
	"fmt"
	"sort"
	"strings"
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
	return fmt.Sprintf("envvar.%s: invalid value %q for environment variable %q", e.Func, e.Value, e.Key)
}

func (e ParseError) Unwrap() error {
	return e.Err
}

// VarSetParseError represents the errors collected while parsing the
// environment variables registered with a [VarSet], keyed by environment
// variable name.
//
// It holds a map and is therefore not comparable: use [GetError], a type
// assertion or errors.As to inspect it, never errors.Is.
type VarSetParseError struct {
	Errors map[string]error
}

// GetError returns the error for the given key if any.
func (e VarSetParseError) GetError(key string) error {
	return e.Errors[key]
}

func (e VarSetParseError) Error() string {
	keys := e.sortedKeys()

	var sb strings.Builder
	sb.WriteString("envvar.VarSet.Parse: parsing failed for: ")
	sb.WriteString(strings.Join(keys, ","))

	for _, key := range keys {
		sb.WriteString("\n")
		sb.WriteString(e.Errors[key].Error())
	}

	return sb.String()
}

func (e VarSetParseError) sortedKeys() []string {
	keys := make([]string, 0, len(e.Errors))

	for key := range e.Errors {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}

func (e VarSetParseError) Unwrap() []error {
	keys := e.sortedKeys()
	errs := make([]error, 0, len(keys))

	for _, key := range keys {
		errs = append(errs, e.Errors[key])
	}

	return errs
}
