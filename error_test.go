package envvar_test

import (
	"errors"
	"testing"

	"github.com/tommarien/envvar"
)

func TestNotSetError(t *testing.T) {
	err := envvar.NotSetError{
		Func: "String",
		Key:  "ENV_VAR",
	}

	want := "envvar.String: missing required environment variable \"ENV_VAR\""
	got := err.Error()

	assertEqual(t, got, want)
}

func TestConversionError(t *testing.T) {
	underlyingErr := errors.New("some error")

	err := envvar.ConversionError{
		Func:  "Int",
		Key:   "ENV_VAR",
		Value: "not an int",
		Err:   underlyingErr,
	}

	want := "envvar.Int: invalid value \"not an int\" for environment variable \"ENV_VAR\": some error"
	got := err.Error()
	assertEqual(t, got, want)

	// verify is (satisfied by unwrap)
	if !errors.Is(err, underlyingErr) {
		t.Fatalf("got error of type %T, want %T", err, underlyingErr)
	}
}
