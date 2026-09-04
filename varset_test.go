package envvar_test

import (
	"errors"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/tommarien/envvar"
)

func assertIsVarSetParseErr(t *testing.T, got error) envvar.VarSetParseError {
	t.Helper()

	parseErr, ok := got.(envvar.VarSetParseError)
	if !ok {
		t.Fatalf("expected VarSetParseError, got %T", got)
	}

	return parseErr
}

func TestStringVar(t *testing.T) {
	t.Run("var set", func(t *testing.T) {
		t.Setenv("ENV_VAR", "value")

		var s envvar.VarSet
		var got string
		envvar.StringVar(&s, &got, "ENV_VAR")

		assertIsNil(t, s.Parse())
		assertEqual(t, got, "value")
	})

	t.Run("var not set", func(t *testing.T) {
		var s envvar.VarSet
		var got string
		envvar.StringVar(&s, &got, "ENV_VAR")

		parseErr := assertIsVarSetParseErr(t, s.Parse())
		assertIsNotSetErr(t, "String", "ENV_VAR", parseErr.GetError("ENV_VAR"))
		assertEqual(t, got, "")
	})
}

func TestStringVarOrDefault(t *testing.T) {
	t.Run("var set", func(t *testing.T) {
		t.Setenv("ENV_VAR", "value")

		var s envvar.VarSet
		var got string
		envvar.StringVarOrDefault(&s, &got, "ENV_VAR", "default")

		assertIsNil(t, s.Parse())
		assertEqual(t, got, "value")
	})

	t.Run("var not set", func(t *testing.T) {
		var s envvar.VarSet
		var got string
		envvar.StringVarOrDefault(&s, &got, "ENV_VAR", "default")

		assertIsNil(t, s.Parse())
		assertEqual(t, got, "default")
	})
}

func TestIntVar(t *testing.T) {
	t.Run("var set", func(t *testing.T) {
		t.Setenv("ENV_VAR", "42")

		var s envvar.VarSet
		var got int
		envvar.IntVar(&s, &got, "ENV_VAR")

		assertIsNil(t, s.Parse())
		assertEqual(t, got, 42)
	})

	t.Run("var set to an invalid value", func(t *testing.T) {
		t.Setenv("ENV_VAR", "not an int")

		var s envvar.VarSet
		var got int
		envvar.IntVar(&s, &got, "ENV_VAR")

		parseErr := assertIsVarSetParseErr(t, s.Parse())
		assertIsParseErr(t, "Int", "ENV_VAR", "not an int", parseErr.GetError("ENV_VAR"))
	})

	t.Run("var not set", func(t *testing.T) {
		var s envvar.VarSet
		var got int
		envvar.IntVar(&s, &got, "ENV_VAR")

		parseErr := assertIsVarSetParseErr(t, s.Parse())
		assertIsNotSetErr(t, "Int", "ENV_VAR", parseErr.GetError("ENV_VAR"))
	})
}

func TestIntVarOrDefault(t *testing.T) {
	t.Run("var set", func(t *testing.T) {
		t.Setenv("ENV_VAR", "42")

		var s envvar.VarSet
		var got int
		envvar.IntVarOrDefault(&s, &got, "ENV_VAR", 21)

		assertIsNil(t, s.Parse())
		assertEqual(t, got, 42)
	})

	t.Run("var set to an invalid value", func(t *testing.T) {
		t.Setenv("ENV_VAR", "not an int")

		var s envvar.VarSet
		var got int
		envvar.IntVarOrDefault(&s, &got, "ENV_VAR", 21)

		parseErr := assertIsVarSetParseErr(t, s.Parse())
		assertIsParseErr(t, "IntOrDefault", "ENV_VAR", "not an int", parseErr.GetError("ENV_VAR"))
	})

	t.Run("var not set", func(t *testing.T) {
		var s envvar.VarSet
		var got int
		envvar.IntVarOrDefault(&s, &got, "ENV_VAR", 21)

		assertIsNil(t, s.Parse())
		assertEqual(t, got, 21)
	})
}

func TestBoolVar(t *testing.T) {
	t.Run("var set", func(t *testing.T) {
		t.Setenv("ENV_VAR", "true")

		var s envvar.VarSet
		var got bool
		envvar.BoolVar(&s, &got, "ENV_VAR")

		assertIsNil(t, s.Parse())
		assertEqual(t, got, true)
	})

	t.Run("var set to an invalid value", func(t *testing.T) {
		t.Setenv("ENV_VAR", "not a bool")

		var s envvar.VarSet
		var got bool
		envvar.BoolVar(&s, &got, "ENV_VAR")

		parseErr := assertIsVarSetParseErr(t, s.Parse())
		assertIsParseErr(t, "Bool", "ENV_VAR", "not a bool", parseErr.GetError("ENV_VAR"))
	})

	t.Run("var not set", func(t *testing.T) {
		var s envvar.VarSet
		var got bool
		envvar.BoolVar(&s, &got, "ENV_VAR")

		parseErr := assertIsVarSetParseErr(t, s.Parse())
		assertIsNotSetErr(t, "Bool", "ENV_VAR", parseErr.GetError("ENV_VAR"))
	})
}

func TestBoolVarOrDefault(t *testing.T) {
	t.Run("var set", func(t *testing.T) {
		t.Setenv("ENV_VAR", "false")

		var s envvar.VarSet
		var got bool
		envvar.BoolVarOrDefault(&s, &got, "ENV_VAR", true)

		assertIsNil(t, s.Parse())
		assertEqual(t, got, false)
	})

	t.Run("var set to an invalid value", func(t *testing.T) {
		t.Setenv("ENV_VAR", "not a bool")

		var s envvar.VarSet
		var got bool
		envvar.BoolVarOrDefault(&s, &got, "ENV_VAR", true)

		parseErr := assertIsVarSetParseErr(t, s.Parse())
		assertIsParseErr(t, "BoolOrDefault", "ENV_VAR", "not a bool", parseErr.GetError("ENV_VAR"))
	})

	t.Run("var not set", func(t *testing.T) {
		var s envvar.VarSet
		var got bool
		envvar.BoolVarOrDefault(&s, &got, "ENV_VAR", true)

		assertIsNil(t, s.Parse())
		assertEqual(t, got, true)
	})
}

func TestUIntVar(t *testing.T) {
	t.Run("var set", func(t *testing.T) {
		t.Setenv("ENV_VAR", "42")

		var s envvar.VarSet
		var got uint
		envvar.UIntVar(&s, &got, "ENV_VAR")

		assertIsNil(t, s.Parse())
		assertEqual(t, got, uint(42))
	})

	t.Run("var set to an invalid value", func(t *testing.T) {
		t.Setenv("ENV_VAR", "-1")

		var s envvar.VarSet
		var got uint
		envvar.UIntVar(&s, &got, "ENV_VAR")

		parseErr := assertIsVarSetParseErr(t, s.Parse())
		assertIsParseErr(t, "UInt", "ENV_VAR", "-1", parseErr.GetError("ENV_VAR"))
	})

	t.Run("var not set", func(t *testing.T) {
		var s envvar.VarSet
		var got uint
		envvar.UIntVar(&s, &got, "ENV_VAR")

		parseErr := assertIsVarSetParseErr(t, s.Parse())
		assertIsNotSetErr(t, "UInt", "ENV_VAR", parseErr.GetError("ENV_VAR"))
	})
}

func TestUIntVarOrDefault(t *testing.T) {
	t.Run("var set", func(t *testing.T) {
		t.Setenv("ENV_VAR", "42")

		var s envvar.VarSet
		var got uint
		envvar.UIntVarOrDefault(&s, &got, "ENV_VAR", 21)

		assertIsNil(t, s.Parse())
		assertEqual(t, got, uint(42))
	})

	t.Run("var set to an invalid value", func(t *testing.T) {
		t.Setenv("ENV_VAR", "-1")

		var s envvar.VarSet
		var got uint
		envvar.UIntVarOrDefault(&s, &got, "ENV_VAR", 21)

		parseErr := assertIsVarSetParseErr(t, s.Parse())
		assertIsParseErr(t, "UIntOrDefault", "ENV_VAR", "-1", parseErr.GetError("ENV_VAR"))
	})

	t.Run("var not set", func(t *testing.T) {
		var s envvar.VarSet
		var got uint
		envvar.UIntVarOrDefault(&s, &got, "ENV_VAR", 21)

		assertIsNil(t, s.Parse())
		assertEqual(t, got, uint(21))
	})
}

func TestFloat64Var(t *testing.T) {
	t.Run("var set", func(t *testing.T) {
		t.Setenv("ENV_VAR", "4.2")

		var s envvar.VarSet
		var got float64
		envvar.Float64Var(&s, &got, "ENV_VAR")

		assertIsNil(t, s.Parse())
		assertEqual(t, got, 4.2)
	})

	t.Run("var set to an invalid value", func(t *testing.T) {
		t.Setenv("ENV_VAR", "not a float")

		var s envvar.VarSet
		var got float64
		envvar.Float64Var(&s, &got, "ENV_VAR")

		parseErr := assertIsVarSetParseErr(t, s.Parse())
		assertIsParseErr(t, "Float64", "ENV_VAR", "not a float", parseErr.GetError("ENV_VAR"))
	})

	t.Run("var not set", func(t *testing.T) {
		var s envvar.VarSet
		var got float64
		envvar.Float64Var(&s, &got, "ENV_VAR")

		parseErr := assertIsVarSetParseErr(t, s.Parse())
		assertIsNotSetErr(t, "Float64", "ENV_VAR", parseErr.GetError("ENV_VAR"))
	})
}

func TestFloat64VarOrDefault(t *testing.T) {
	t.Run("var set", func(t *testing.T) {
		t.Setenv("ENV_VAR", "4.2")

		var s envvar.VarSet
		var got float64
		envvar.Float64VarOrDefault(&s, &got, "ENV_VAR", 2.1)

		assertIsNil(t, s.Parse())
		assertEqual(t, got, 4.2)
	})

	t.Run("var set to an invalid value", func(t *testing.T) {
		t.Setenv("ENV_VAR", "not a float")

		var s envvar.VarSet
		var got float64
		envvar.Float64VarOrDefault(&s, &got, "ENV_VAR", 2.1)

		parseErr := assertIsVarSetParseErr(t, s.Parse())
		assertIsParseErr(t, "Float64OrDefault", "ENV_VAR", "not a float", parseErr.GetError("ENV_VAR"))
	})

	t.Run("var not set", func(t *testing.T) {
		var s envvar.VarSet
		var got float64
		envvar.Float64VarOrDefault(&s, &got, "ENV_VAR", 2.1)

		assertIsNil(t, s.Parse())
		assertEqual(t, got, 2.1)
	})
}

func TestDurationVar(t *testing.T) {
	t.Run("var set", func(t *testing.T) {
		t.Setenv("ENV_VAR", "5m")

		var s envvar.VarSet
		var got time.Duration
		envvar.DurationVar(&s, &got, "ENV_VAR")

		assertIsNil(t, s.Parse())
		assertEqual(t, got, 5*time.Minute)
	})

	t.Run("var set to an invalid value", func(t *testing.T) {
		t.Setenv("ENV_VAR", "not a duration")

		var s envvar.VarSet
		var got time.Duration
		envvar.DurationVar(&s, &got, "ENV_VAR")

		parseErr := assertIsVarSetParseErr(t, s.Parse())
		assertIsParseErr(t, "Duration", "ENV_VAR", "not a duration", parseErr.GetError("ENV_VAR"))
	})

	t.Run("var not set", func(t *testing.T) {
		var s envvar.VarSet
		var got time.Duration
		envvar.DurationVar(&s, &got, "ENV_VAR")

		parseErr := assertIsVarSetParseErr(t, s.Parse())
		assertIsNotSetErr(t, "Duration", "ENV_VAR", parseErr.GetError("ENV_VAR"))
	})
}

func TestDurationVarOrDefault(t *testing.T) {
	t.Run("var set", func(t *testing.T) {
		t.Setenv("ENV_VAR", "5m")

		var s envvar.VarSet
		var got time.Duration
		envvar.DurationVarOrDefault(&s, &got, "ENV_VAR", 10*time.Minute)

		assertIsNil(t, s.Parse())
		assertEqual(t, got, 5*time.Minute)
	})

	t.Run("var set to an invalid value", func(t *testing.T) {
		t.Setenv("ENV_VAR", "not a duration")

		var s envvar.VarSet
		var got time.Duration
		envvar.DurationVarOrDefault(&s, &got, "ENV_VAR", 10*time.Minute)

		parseErr := assertIsVarSetParseErr(t, s.Parse())
		assertIsParseErr(t, "DurationOrDefault", "ENV_VAR", "not a duration", parseErr.GetError("ENV_VAR"))
	})

	t.Run("var not set", func(t *testing.T) {
		var s envvar.VarSet
		var got time.Duration
		envvar.DurationVarOrDefault(&s, &got, "ENV_VAR", 10*time.Minute)

		assertIsNil(t, s.Parse())
		assertEqual(t, got, 10*time.Minute)
	})
}

// parseLevel is a caller supplied [envvar.ParseFunc], proving the escape
// hatch works for a type the package has no accessor for.
func parseLevel(value string) (level, error) {
	switch strings.ToLower(value) {
	case "debug":
		return levelDebug, nil
	case "info":
		return levelInfo, nil
	default:
		return 0, errors.New("unknown level")
	}
}

type level int

const (
	levelDebug level = iota + 1
	levelInfo
)

func TestFuncVar(t *testing.T) {
	t.Run("var set", func(t *testing.T) {
		t.Setenv("ENV_VAR", "info")

		var s envvar.VarSet
		var got level
		envvar.FuncVar(&s, &got, "ENV_VAR", "Level", parseLevel)

		assertIsNil(t, s.Parse())
		assertEqual(t, got, levelInfo)
	})

	t.Run("var set to an invalid value", func(t *testing.T) {
		t.Setenv("ENV_VAR", "nope")

		var s envvar.VarSet
		var got level
		envvar.FuncVar(&s, &got, "ENV_VAR", "Level", parseLevel)

		parseErr := assertIsVarSetParseErr(t, s.Parse())
		assertIsParseErr(t, "Level", "ENV_VAR", "nope", parseErr.GetError("ENV_VAR"))
	})

	t.Run("var not set", func(t *testing.T) {
		var s envvar.VarSet
		var got level
		envvar.FuncVar(&s, &got, "ENV_VAR", "Level", parseLevel)

		parseErr := assertIsVarSetParseErr(t, s.Parse())
		assertIsNotSetErr(t, "Level", "ENV_VAR", parseErr.GetError("ENV_VAR"))
	})
}

func TestFuncVarOrDefault(t *testing.T) {
	t.Run("var set", func(t *testing.T) {
		t.Setenv("ENV_VAR", "debug")

		var s envvar.VarSet
		var got level
		envvar.FuncVarOrDefault(&s, &got, "ENV_VAR", levelInfo, "LevelOrDefault", parseLevel)

		assertIsNil(t, s.Parse())
		assertEqual(t, got, levelDebug)
	})

	t.Run("var set to an invalid value", func(t *testing.T) {
		t.Setenv("ENV_VAR", "nope")

		var s envvar.VarSet
		var got level
		envvar.FuncVarOrDefault(&s, &got, "ENV_VAR", levelInfo, "LevelOrDefault", parseLevel)

		parseErr := assertIsVarSetParseErr(t, s.Parse())
		assertIsParseErr(t, "LevelOrDefault", "ENV_VAR", "nope", parseErr.GetError("ENV_VAR"))
	})

	t.Run("var not set", func(t *testing.T) {
		var s envvar.VarSet
		var got level
		envvar.FuncVarOrDefault(&s, &got, "ENV_VAR", levelInfo, "LevelOrDefault", parseLevel)

		assertIsNil(t, s.Parse())
		assertEqual(t, got, levelInfo)
	})
}

func TestVarSetParse(t *testing.T) {
	t.Run("returns nil when nothing is registered", func(t *testing.T) {
		var s envvar.VarSet

		assertIsNil(t, s.Parse())
	})

	t.Run("assigns every registered var", func(t *testing.T) {
		t.Setenv("NAME", "envvar")
		t.Setenv("PORT", "8080")

		var s envvar.VarSet
		var name, region string
		var port int
		var timeout time.Duration

		envvar.StringVar(&s, &name, "NAME")
		envvar.StringVarOrDefault(&s, &region, "REGION", "eu-west-1")
		envvar.IntVar(&s, &port, "PORT")
		envvar.DurationVarOrDefault(&s, &timeout, "TIMEOUT", 10*time.Minute)

		assertIsNil(t, s.Parse())
		assertEqual(t, name, "envvar")
		assertEqual(t, region, "eu-west-1")
		assertEqual(t, port, 8080)
		assertEqual(t, timeout, 10*time.Minute)
	})

	t.Run("aggregates the errors of every failing key", func(t *testing.T) {
		t.Setenv("PORT", "not an int")

		var s envvar.VarSet
		var name string
		var port int
		var enabled bool

		envvar.StringVar(&s, &name, "NAME")
		envvar.IntVar(&s, &port, "PORT")
		envvar.BoolVar(&s, &enabled, "ENABLED")

		parseErr := assertIsVarSetParseErr(t, s.Parse())

		assertIsNotSetErr(t, "String", "NAME", parseErr.GetError("NAME"))
		assertIsParseErr(t, "Int", "PORT", "not an int", parseErr.GetError("PORT"))
		assertIsNotSetErr(t, "Bool", "ENABLED", parseErr.GetError("ENABLED"))
		assertEqual(t, len(parseErr.Errors), 3)
	})

	t.Run("runs the actions in registration order", func(t *testing.T) {
		t.Setenv("FIRST", "1")
		t.Setenv("SECOND", "2")
		t.Setenv("THIRD", "3")

		var s envvar.VarSet
		var order []string
		var first, second, third int

		envvar.FuncVar(&s, &first, "FIRST", "First", func(v string) (int, error) {
			order = append(order, "FIRST")
			return strconv.Atoi(v)
		})
		envvar.FuncVar(&s, &second, "SECOND", "Second", func(v string) (int, error) {
			order = append(order, "SECOND")
			return strconv.Atoi(v)
		})
		envvar.FuncVar(&s, &third, "THIRD", "Third", func(v string) (int, error) {
			order = append(order, "THIRD")
			return strconv.Atoi(v)
		})

		assertIsNil(t, s.Parse())
		assertEqual(t, strings.Join(order, ","), "FIRST,SECOND,THIRD")
	})

	t.Run("runs every registration of a duplicate key", func(t *testing.T) {
		t.Setenv("ENV_VAR", "42")

		var s envvar.VarSet
		var first, second int

		envvar.IntVar(&s, &first, "ENV_VAR")
		envvar.IntVar(&s, &second, "ENV_VAR")

		assertIsNil(t, s.Parse())
		assertEqual(t, first, 42)
		assertEqual(t, second, 42)
	})

	t.Run("keeps the last error of a duplicate key", func(t *testing.T) {
		var s envvar.VarSet
		var first int
		var second bool

		envvar.IntVar(&s, &first, "ENV_VAR")
		envvar.BoolVar(&s, &second, "ENV_VAR")

		parseErr := assertIsVarSetParseErr(t, s.Parse())
		assertIsNotSetErr(t, "Bool", "ENV_VAR", parseErr.GetError("ENV_VAR"))
		assertEqual(t, len(parseErr.Errors), 1)
	})
}

func TestVarSetNil(t *testing.T) {
	defer func() {
		got := recover()
		if got == nil {
			t.Fatal("expected panic, got none")
		}

		message, ok := got.(string)
		if !ok {
			t.Fatalf("expected a string panic, got %T", got)
		}

		assertEqual(t, message, "envvar: nil *VarSet")
	}()

	var value string
	envvar.StringVar(nil, &value, "ENV_VAR")
}
