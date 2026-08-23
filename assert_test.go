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

func assertIsNotSetErr(t *testing.T, funcName, key string, got error) {
	t.Helper()

	notSetErr, ok := got.(envvar.NotSetError)
	if !ok {
		t.Fatalf("expected NotSetError, got %T", got)
	}

	if notSetErr.Func != funcName {
		t.Errorf("expected NotSetError with func %q, got %q", funcName, notSetErr.Func)
	}

	if notSetErr.Key != key {
		t.Errorf("expected NotSetError with key %q, got %q", key, notSetErr.Key)
	}
}

func assertIsParseErr(t *testing.T, funcName, key, value string, got error) {
	t.Helper()

	parseErr, ok := got.(envvar.ParseError)
	if !ok {
		t.Fatalf("expected ParseError, got %T", got)
	}

	if parseErr.Func != funcName {
		t.Errorf("expected ParseError with func %q, got %q", funcName, parseErr.Func)
	}

	if parseErr.Key != key {
		t.Errorf("expected ParseError with key %q, got %q", key, parseErr.Key)
	}

	if parseErr.Value != value {
		t.Errorf("expected ParseError with value %q, got %q", value, parseErr.Value)
	}
}
