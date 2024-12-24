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

	assertEqual(t, err.Error(), "envvar.String: missing required environment variable \"ENV_VAR\"")
}

func TestParseError(t *testing.T) {
	underlyingErr := errors.New("some error")

	err := envvar.ParseError{
		Func:  "Int",
		Key:   "ENV_VAR",
		Value: "not an int",
		Err:   underlyingErr,
	}

	assertEqual(t, err.Error(), "envvar.Int: invalid value \"not an int\" for environment variable \"ENV_VAR\": some error")

	// unwrap
	assertEqual(t, err.Unwrap(), underlyingErr)

	// verify is (satisfied by unwrap)
	var wantErr error = underlyingErr
	if !errors.Is(err, wantErr) {
		t.Fatalf("got error of type %T, want %T", err, wantErr)
	}
}

func TestString(t *testing.T) {
	t.Run("var set", func(t *testing.T) {
		t.Setenv("ENV_VAR", "value")

		got, err := envvar.String("ENV_VAR")
		assertIsNil(t, err)
		assertEqual(t, got, "value")
	})

	t.Run("var set to zero value", func(t *testing.T) {
		t.Setenv("ENV_VAR", "")

		_, err := envvar.String("ENV_VAR")
		assertIsNil(t, err)
	})

	t.Run("var not set", func(t *testing.T) {
		_, err := envvar.String("ENV_VAR")
		assertIsNotSetErr(t, "ENV_VAR", err)
	})
}

func TestStringOrDefault(t *testing.T) {
	t.Run("var set", func(t *testing.T) {
		t.Setenv("ENV_VAR", "value")

		got := envvar.StringOrDefault("ENV_VAR", "default")
		assertEqual(t, got, "value")
	})

	t.Run("var set to zero value", func(t *testing.T) {
		t.Setenv("ENV_VAR", "")

		got := envvar.StringOrDefault("ENV_VAR", "default")
		assertEqual(t, got, "")
	})

	t.Run("var not set", func(t *testing.T) {
		got := envvar.StringOrDefault("ENV_VAR", "default")
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
}

func TestIntOrDefault(t *testing.T) {
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
			name: "var set to negative value",
			env:  "-10",
			want: -10,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("ENV_VAR", tt.env)

			got, err := envvar.IntOrDefault("ENV_VAR", -99)
			assertIsNil(t, err)
			assertEqual(t, got, tt.want)
		})
	}

	t.Run("var set to invalid integer", func(t *testing.T) {
		t.Setenv("ENV_VAR", "1.1")

		_, err := envvar.IntOrDefault("ENV_VAR", -10)
		var wantErr envvar.ParseError
		if !errors.As(err, &wantErr) {
			t.Fatalf("got error of type %T, want %T", err, wantErr)
		}
		assertEqual(t, wantErr.Key, "ENV_VAR")
		assertEqual(t, wantErr.Value, "1.1")
		assertEqual(t, wantErr.Func, "IntOrDefault")
	})

	t.Run("var set to zero value", func(t *testing.T) {
		t.Setenv("ENV_VAR", "0")

		got, err := envvar.IntOrDefault("ENV_VAR", -99)
		assertIsNil(t, err)
		assertEqual(t, got, 0)
	})

	t.Run("var not set returns the default value", func(t *testing.T) {
		got, err := envvar.IntOrDefault("ENV_VAR", 10)
		assertIsNil(t, err)
		assertEqual(t, got, 10)
	})
}
