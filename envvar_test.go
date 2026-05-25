package envvar_test

import (
	"fmt"
	"testing"

	"github.com/tommarien/envvar"
)

func TestString(t *testing.T) {
	t.Run("var set", func(t *testing.T) {
		t.Setenv("ENV_VAR", "value")

		got, err := envvar.String("ENV_VAR")
		assertIsNil(t, err)
		assertEqual(t, got, "value")
	})

	t.Run("var set to zero value", func(t *testing.T) {
		t.Setenv("ENV_VAR", "")

		got, err := envvar.String("ENV_VAR")
		assertIsNil(t, err)
		assertEqual(t, got, "")
	})

	t.Run("var not set", func(t *testing.T) {
		_, err := envvar.String("ENV_VAR")
		assertIsNotSetErr(t, "ENV_VAR", err)
	})

	t.Run("var not set with default option", func(t *testing.T) {
		got, err := envvar.String("ENV_VAR", envvar.WithDefault("default"))
		assertIsNil(t, err)
		assertEqual(t, got, "default")
	})
}

func TestInt(t *testing.T) {
	cases := []struct {
		name string
		env  string
		want int
	}{
		{
			name: "var set",
			env:  "99",
			want: 99,
		},
		{
			name: "var set to zero value",
			env:  "0",
			want: 0,
		},
		{
			name: "var set to negative value",
			env:  "-1",
			want: -1,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("ENV_VAR", tt.env)

			got, err := envvar.Int("ENV_VAR")
			assertIsNil(t, err)
			assertEqual(t, got, tt.want)
		})
	}

	t.Run("var set to invalid value", func(t *testing.T) {
		t.Setenv("ENV_VAR", "not an int")

		_, err := envvar.Int("ENV_VAR")

		convErr, ok := err.(envvar.ParseError)
		if !ok {
			t.Fatalf("got error of type %T, want ParseError", err)
		}

		assertEqual(t, convErr.Key, "ENV_VAR")
		assertEqual(t, convErr.Value, "not an int")
	})

	t.Run("var not set", func(t *testing.T) {
		_, err := envvar.Int("ENV_VAR")
		assertIsNotSetErr(t, "ENV_VAR", err)
		assertEqual(t, err.Error(), "envvar.Int: missing required environment variable \"ENV_VAR\"")
	})

	t.Run("var not set with default option", func(t *testing.T) {
		got, err := envvar.Int("ENV_VAR", envvar.WithDefault(30))
		assertIsNil(t, err)
		assertEqual(t, got, 30)
	})
}

func TestBool(t *testing.T) {
	trueCases := []string{"1", "true", "t", "True", "T"}
	falseCases := []string{"0", "false", "f", "False", "F"}

	for _, trueString := range trueCases {
		t.Run(fmt.Sprintf("var set to %s returns true", trueString), func(t *testing.T) {
			t.Setenv("ENV_VAR", trueString)

			got, err := envvar.Bool("ENV_VAR")
			assertIsNil(t, err)
			assertEqual(t, got, true)
		})
	}

	for _, falseString := range falseCases {
		t.Run(fmt.Sprintf("var set to %s returns false", falseString), func(t *testing.T) {
			t.Setenv("ENV_VAR", falseString)

			got, err := envvar.Bool("ENV_VAR")
			assertIsNil(t, err)
			assertEqual(t, got, false)
		})
	}

	t.Run("var set to invalid value", func(t *testing.T) {
		t.Setenv("ENV_VAR", "not a bool")

		_, err := envvar.Bool("ENV_VAR")

		convErr, ok := err.(envvar.ParseError)
		if !ok {
			t.Fatalf("got error of type %T, want ParseError", err)
		}

		assertEqual(t, convErr.Key, "ENV_VAR")
		assertEqual(t, convErr.Value, "not a bool")
	})

	t.Run("var not set", func(t *testing.T) {
		_, err := envvar.Bool("ENV_VAR")
		assertIsNotSetErr(t, "ENV_VAR", err)
		assertEqual(t, err.Error(), "envvar.Bool: missing required environment variable \"ENV_VAR\"")
	})

	t.Run("var not set with default option", func(t *testing.T) {
		got, err := envvar.Bool("ENV_VAR", envvar.WithDefault(false))
		assertIsNil(t, err)
		assertEqual(t, got, false)
	})
}
