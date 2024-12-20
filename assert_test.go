package envvar_test

import (
	"errors"
	"testing"

	"github.com/tommarien/envvar"
)

func assertIsNil(t *testing.T, got any) {
	t.Helper()
	if got != nil {
		t.Errorf("expected nil, got %v", got)
	}
}

func assertEqual(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func assertIsNotSetErr(t *testing.T, key string, got error) {
	t.Helper()

	var err envvar.NotSetError
	if !errors.As(got, &err) {
		t.Errorf("expected NotSetError, got %T", got)
		return
	}

	if err.Key != key {
		t.Errorf("expected NotSetError with key %q, got %q", key, err.Key)
	}
}
