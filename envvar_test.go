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
	for _, tc := range []struct {
		value string
		want  int
	}{
		{"-1", -1},
		{"0", 0},
		{"1", 1},
		{"10", 10},
	} {
		t.Run("var set to "+tc.value, func(t *testing.T) {
			t.Setenv("ENV_VAR", tc.value)

			got, err := envvar.Int("ENV_VAR")
			assertIsNil(t, err)
			assertEqual(t, got, tc.want)
		})
	}

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
	for _, tc := range []struct {
		value string
		want  int
	}{
		{"-1", -1},
		{"0", 0},
		{"1", 1},
		{"10", 10},
	} {
		t.Run("var set to "+tc.value, func(t *testing.T) {
			t.Setenv("ENV_VAR", tc.value)

			got, err := envvar.IntOrDefault("ENV_VAR", 42)
			assertIsNil(t, err)
			assertEqual(t, got, tc.want)
		})
	}

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

func TestBool(t *testing.T) {
	for _, tc := range []struct {
		value string
		want  bool
	}{
		{"true", true},
		{"false", false},
		{"1", true},
		{"0", false},
	} {
		t.Run("var set to "+tc.value, func(t *testing.T) {
			t.Setenv("ENV_VAR", tc.value)

			got, err := envvar.Bool("ENV_VAR")
			assertIsNil(t, err)
			assertEqual(t, got, tc.want)
		})
	}

	t.Run("var not set", func(t *testing.T) {
		_, err := envvar.Bool("ENV_VAR")
		assertIsNotSetErr(t, "Bool", "ENV_VAR", err)
	})

	t.Run("var set to invalid value", func(t *testing.T) {
		t.Setenv("ENV_VAR", "A")

		_, err := envvar.Bool("ENV_VAR")
		assertIsParseErr(t, "Bool", "ENV_VAR", "A", err)
	})
}

func TestBoolOrDefault(t *testing.T) {
	for _, tc := range []struct {
		value string
		want  bool
	}{
		{"true", true},
		{"false", false},
		{"1", true},
		{"0", false},
	} {
		t.Run("var set to "+tc.value, func(t *testing.T) {
			t.Setenv("ENV_VAR", tc.value)

			got, err := envvar.BoolOrDefault("ENV_VAR", true)
			assertIsNil(t, err)
			assertEqual(t, got, tc.want)
		})
	}

	t.Run("var not set", func(t *testing.T) {
		got, err := envvar.BoolOrDefault("ENV_VAR", true)
		assertIsNil(t, err)
		assertEqual(t, got, true)
	})

	t.Run("var set to invalid value", func(t *testing.T) {
		t.Setenv("ENV_VAR", "A")

		_, err := envvar.BoolOrDefault("ENV_VAR", true)
		assertIsParseErr(t, "BoolOrDefault", "ENV_VAR", "A", err)
	})
}
