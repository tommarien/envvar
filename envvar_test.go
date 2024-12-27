package envvar_test

import (
	"errors"
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

		var wantErr envvar.ParseError
		if !errors.As(err, &wantErr) {
			t.Fatalf("got error of type %T, want %T", err, wantErr)
		}

		assertEqual(t, wantErr.Key, "ENV_VAR")
		assertEqual(t, wantErr.Value, "not an int")
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
