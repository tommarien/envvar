package envvar

import (
	"errors"
	"fmt"
	"slices"
	"strings"
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
	Errors map[string]error
}

// GetError returns the error for the given key if any.
func (p ParserError) GetError(key string) error {
	return p.Errors[key]
}

func (p ParserError) Error() string {
	var sb strings.Builder
	sb.WriteString("Parser.Parse: failed for the following envvars: ")

	keys := p.sortedKeys()
	errs := make([]error, 0, len(keys))

	for _, key := range keys {
		errs = append(errs, p.Errors[key])
	}

	sb.WriteString(strings.Join(keys, ","))
	sb.WriteString("\n")

	sb.WriteString(errors.Join(errs...).Error())

	return sb.String()
}

func (p ParserError) sortedKeys() []string {
	keys := make([]string, 0, len(p.Errors))

	for key := range p.Errors {
		keys = append(keys, key)
	}

	slices.Sort(keys)

	return keys
}

func (p ParserError) Unwrap() []error {
	keys := p.sortedKeys()
	errs := make([]error, 0, len(keys))

	for _, key := range keys {
		errs = append(errs, p.Errors[key])
	}
	return errs
}
