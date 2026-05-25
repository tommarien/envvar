package envvar_test

import (
	"testing"

	"github.com/tommarien/envvar"
)

func TestParserString(t *testing.T) {
	t.Run("sets the value to env var when its set", func(t *testing.T) {
		t.Setenv("ENV_VAR", "value")

		var value string
		parser := envvar.NewGroup()

		parser.StringVar(&value, "ENV_VAR")

		err := parser.Parse()
		assertIsNil(t, err)
		assertEqual(t, value, "value")
	})

	t.Run("sets the value to default when specified", func(t *testing.T) {
		parser := envvar.NewGroup()

		var value string
		parser.StringVar(&value, "ENV_VAR", envvar.WithDefault("default"))

		err := parser.Parse()
		assertIsNil(t, err)
		assertEqual(t, value, "default")
	})

	t.Run("returns a parser error when the value is unset", func(t *testing.T) {
		parser := envvar.NewGroup()

		var value string
		parser.StringVar(&value, "ENV_VAR")

		err := parser.Parse()
		if err == nil {
			t.Fatal("expected error got nil")
		}

		pErr, ok := err.(envvar.GroupParseError)
		if !ok {
			t.Fatalf("expected ParserError got %T", err)
		}

		errOnEnvVar := pErr.GetError("ENV_VAR")
		if _, ok := errOnEnvVar.(envvar.NotSetError); !ok {
			t.Fatalf("expected NotSetError on ENV_VAR got %T", errOnEnvVar)
		}
	})
}

func TestIntVar(t *testing.T) {
	t.Run("sets the value to the env var when set", func(t *testing.T) {
		t.Setenv("ENV_VAR", "42")

		var value int
		parser := envvar.NewGroup()

		parser.IntVar(&value, "ENV_VAR")

		err := parser.Parse()
		assertIsNil(t, err)
		assertEqual(t, value, 42)
	})

	t.Run("sets the value to default when specified", func(t *testing.T) {
		parser := envvar.NewGroup()

		var value int
		parser.IntVar(&value, "ENV_VAR", envvar.WithDefault(42))

		err := parser.Parse()
		assertIsNil(t, err)
		assertEqual(t, value, 42)
	})

	t.Run("returns a parser error when the value is unset", func(t *testing.T) {
		parser := envvar.NewGroup()

		var value int
		parser.IntVar(&value, "ENV_VAR")

		err := parser.Parse()
		if err == nil {
			t.Fatal("expected error got nil")
		}

		pErr, ok := err.(envvar.GroupParseError)
		if !ok {
			t.Fatalf("expected ParserError got %T", err)
		}

		errOnEnvVar := pErr.GetError("ENV_VAR")
		if _, ok := errOnEnvVar.(envvar.NotSetError); !ok {
			t.Fatalf("expected NotSetError on ENV_VAR got %T", errOnEnvVar)
		}
	})
}

func TestBoolVar(t *testing.T) {
	t.Run("sets the value to the env var when set", func(t *testing.T) {
		t.Setenv("ENV_VAR", "true")

		var value bool
		parser := envvar.NewGroup()

		parser.BoolVar(&value, "ENV_VAR")

		err := parser.Parse()
		assertIsNil(t, err)
		assertEqual(t, value, true)
	})

	t.Run("sets the value to default when specified", func(t *testing.T) {
		parser := envvar.NewGroup()

		var value bool
		parser.BoolVar(&value, "ENV_VAR", envvar.WithDefault(true))

		err := parser.Parse()
		assertIsNil(t, err)
		assertEqual(t, value, true)
	})

	t.Run("returns a parser error when the value is unset", func(t *testing.T) {
		parser := envvar.NewGroup()

		var value bool
		parser.BoolVar(&value, "ENV_VAR")

		err := parser.Parse()
		if err == nil {
			t.Fatal("expected error got nil")
		}

		pErr, ok := err.(envvar.GroupParseError)
		if !ok {
			t.Fatalf("expected ParserError got %T", err)
		}

		errOnEnvVar := pErr.GetError("ENV_VAR")
		if _, ok := errOnEnvVar.(envvar.NotSetError); !ok {
			t.Fatalf("expected NotSetError on ENV_VAR got %T", errOnEnvVar)
		}
	})
}
