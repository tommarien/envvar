package envvar_test

import (
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
		assertIsNotSetErr(t, "String", "ENV_VAR", err)
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
	t.Run("var set", func(t *testing.T) {
		t.Setenv("ENV_VAR", "100")

		got, err := envvar.Int("ENV_VAR")
		assertIsNil(t, err)
		assertEqual(t, got, 100)
	})

	t.Run("var set to zero value", func(t *testing.T) {
		t.Setenv("ENV_VAR", "0")

		got, err := envvar.Int("ENV_VAR")
		assertIsNil(t, err)
		assertEqual(t, got, 0)
	})

	t.Run("var not set", func(t *testing.T) {
		_, err := envvar.Int("ENV_VAR")
		assertIsNotSetErr(t, "Int", "ENV_VAR", err)
	})

	t.Run("var set to invalid value", func(t *testing.T) {
		t.Setenv("ENV_VAR", "A")

		_, err := envvar.Int("ENV_VAR")
		assertIsParseErr(t, "Int", "ENV_VAR", "A", err)
	})
}

func TestIntOrDefault(t *testing.T) {
	t.Run("var set", func(t *testing.T) {
		t.Setenv("ENV_VAR", "100")

		got, err := envvar.IntOrDefault("ENV_VAR", 42)
		assertIsNil(t, err)
		assertEqual(t, got, 100)
	})

	t.Run("var set to zero value", func(t *testing.T) {
		t.Setenv("ENV_VAR", "0")

		got, err := envvar.IntOrDefault("ENV_VAR", 42)
		assertIsNil(t, err)
		assertEqual(t, got, 0)
	})

	t.Run("var not set", func(t *testing.T) {
		got, err := envvar.IntOrDefault("ENV_VAR", 42)
		assertIsNil(t, err)
		assertEqual(t, got, 42)
	})

	t.Run("var set to invalid value", func(t *testing.T) {
		t.Setenv("ENV_VAR", "A")

		_, err := envvar.IntOrDefault("ENV_VAR", 42)
		assertIsParseErr(t, "IntOrDefault", "ENV_VAR", "A", err)
	})
}
