package envvar_test

import (
	"errors"
	"fmt"
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

func TestParseError(t *testing.T) {
	underlyingErr := errors.New("some error")

	err := envvar.ParseError{
		Func:  "Int",
		Key:   "ENV_VAR",
		Value: "not an int",
		Err:   underlyingErr,
	}

	want := "envvar.Int: invalid value \"not an int\" for environment variable \"ENV_VAR\""
	got := err.Error()
	assertEqual(t, got, want)

	// verify is (satisfied by unwrap)
	if !errors.Is(err, underlyingErr) {
		t.Fatalf("got error of type %T, want %T", err, underlyingErr)
	}
}

func TestVarSetParseError(t *testing.T) {
	underlyingErr := errors.New("some error")
	underlyingErr2 := errors.New("some other error")

	key1 := "key1"
	key2 := "key2"

	errorMap := make(map[string]error, 2)
	errorMap[key2] = underlyingErr2
	errorMap[key1] = underlyingErr

	err := envvar.VarSetParseError{
		Errors: errorMap,
	}

	t.Run("geterror returns the error on the given key if any", func(t *testing.T) {
		got := err.GetError(key2)
		if !errors.Is(got, underlyingErr2) {
			t.Fatalf("expected %v got %v", underlyingErr2, got)
		}

		unexisting := err.GetError("somekey")
		assertIsNil(t, unexisting)
	})

	t.Run("unwrap returns the errors sorted by key", func(t *testing.T) {
		got := err.Unwrap()

		if len(got) != 2 {
			t.Fatalf("expected 2 errors got %d", len(got))
		}

		if !errors.Is(got[0], underlyingErr) {
			t.Fatalf("expected %v got %v", underlyingErr, got[0])
		}

		if !errors.Is(got[1], underlyingErr2) {
			t.Fatalf("expected %v got %v", underlyingErr2, got[1])
		}
	})

	t.Run("error returns the expected string", func(t *testing.T) {
		got := err.Error()
		assertEqual(
			t,
			got,
			fmt.Sprintf(
				"envvar.VarSet.Parse: parsing failed for: key1,key2\n%v\n%v",
				underlyingErr,
				underlyingErr2,
			),
		)
	})
}
