package envvar_test

import (
	"testing"

	"github.com/tommarien/envvar"
)

func assertIsNil(t *testing.T, got any) {
	t.Helper()
	if got != nil {
		t.Errorf("expected nil, got %v", got)
	}
}

func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("expected %v, got %v", want, got)
	}
}

func assertIsNotSetErr(t *testing.T, key string, got error) {
	t.Helper()

	notSetErr, ok := got.(envvar.NotSetError)
	if !ok {
		t.Fatalf("expected NotSetError, got %T", got)
	}

	if notSetErr.Key != key {
		t.Errorf("expected NotSetError with key %q, got %q", key, notSetErr.Key)
	}
}
