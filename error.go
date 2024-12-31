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

// GroupParseError represents an error when there are errors
// parsing the environment variables using [Group].
type GroupParseError struct {
	Errors map[string]error
}

// GetError returns the error for the given key if any.
func (p GroupParseError) GetError(key string) error {
	return p.Errors[key]
}

func (p GroupParseError) Error() string {
	var sb strings.Builder
	sb.WriteString("Group.Parse: parsing failed for: ")

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

func (p GroupParseError) sortedKeys() []string {
	keys := make([]string, 0, len(p.Errors))

	for key := range p.Errors {
		keys = append(keys, key)
	}

	slices.Sort(keys)

	return keys
}

func (p GroupParseError) Unwrap() []error {
	keys := p.sortedKeys()
	errs := make([]error, 0, len(keys))

	for _, key := range keys {
		errs = append(errs, p.Errors[key])
	}
	return errs
}
