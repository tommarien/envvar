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
