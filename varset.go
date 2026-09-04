package envvar

import "time"

type action struct {
	key string
	run func() error
}

// VarSet collects deferred environment variable lookups, registered through
// the Var functions in this package and run together by [VarSet.Parse] in
// registration order.
//
// The zero value is ready for use. A VarSet must not be copied after the
// first registration, and registering is not safe for concurrent use.
type VarSet struct {
	actions []action
}

func register(s *VarSet, key string, run func() error) {
	if s == nil {
		panic("envvar: nil *VarSet")
	}

	s.actions = append(s.actions, action{key: key, run: run})
}

// varFunc registers key with s using an accessor taking the key. It is
// deliberately unexported: its func(key string) (T, error) shape is
// indistinguishable from [ParseFunc], which takes a value, so FuncVar and
// FuncVarOrDefault expose a [ParseFunc] instead.
func varFunc[T any](s *VarSet, p *T, key string, get func(key string) (T, error)) {
	register(s, key, func() error {
		value, err := get(key)
		if err != nil {
			return err
		}

		*p = value

		return nil
	})
}

// FuncVar registers key with s, mirroring [Func]. When [VarSet.Parse] is
// called, Func(key, funcName, parse) is invoked and its result assigned to
// *p, leaving *p untouched on failure. A failure is collected into the error
// returned by [VarSet.Parse].
func FuncVar[T any](s *VarSet, p *T, key, funcName string, parse ParseFunc[T]) {
	varFunc(s, p, key, func(key string) (T, error) { return Func(key, funcName, parse) })
}

// FuncVarOrDefault registers key with s, mirroring [FuncOrDefault]. When
// [VarSet.Parse] is called, FuncOrDefault(key, defaultValue, funcName, parse)
// is invoked and its result assigned to *p, leaving *p untouched on failure.
// A failure is collected into the error returned by [VarSet.Parse].
func FuncVarOrDefault[T any](s *VarSet, p *T, key string, defaultValue T, funcName string, parse ParseFunc[T]) {
	varFunc(s, p, key, func(key string) (T, error) { return FuncOrDefault(key, defaultValue, funcName, parse) })
}

// StringVar registers key with s, assigning [String]'s result to *p.
func StringVar(s *VarSet, p *string, key string) {
	varFunc(s, p, key, String)
}

// StringVarOrDefault registers key with s, assigning [StringOrDefault]'s
// result to *p. Unlike the other VarOrDefault functions it can never fail,
// as StringOrDefault has no error case.
func StringVarOrDefault(s *VarSet, p *string, key, defaultValue string) {
	varFunc(s, p, key, func(key string) (string, error) {
		return StringOrDefault(key, defaultValue), nil
	})
}

// IntVar registers key with s, assigning [Int]'s result to *p.
func IntVar(s *VarSet, p *int, key string) {
	varFunc(s, p, key, Int)
}

// IntVarOrDefault registers key with s, assigning [IntOrDefault]'s result to *p.
func IntVarOrDefault(s *VarSet, p *int, key string, defaultValue int) {
	varFunc(s, p, key, func(key string) (int, error) { return IntOrDefault(key, defaultValue) })
}

// BoolVar registers key with s, assigning [Bool]'s result to *p.
func BoolVar(s *VarSet, p *bool, key string) {
	varFunc(s, p, key, Bool)
}

// BoolVarOrDefault registers key with s, assigning [BoolOrDefault]'s result to *p.
func BoolVarOrDefault(s *VarSet, p *bool, key string, defaultValue bool) {
	varFunc(s, p, key, func(key string) (bool, error) { return BoolOrDefault(key, defaultValue) })
}

// UIntVar registers key with s, assigning [UInt]'s result to *p.
func UIntVar(s *VarSet, p *uint, key string) {
	varFunc(s, p, key, UInt)
}

// UIntVarOrDefault registers key with s, assigning [UIntOrDefault]'s result to *p.
func UIntVarOrDefault(s *VarSet, p *uint, key string, defaultValue uint) {
	varFunc(s, p, key, func(key string) (uint, error) { return UIntOrDefault(key, defaultValue) })
}

// Float64Var registers key with s, assigning [Float64]'s result to *p.
func Float64Var(s *VarSet, p *float64, key string) {
	varFunc(s, p, key, Float64)
}

// Float64VarOrDefault registers key with s, assigning [Float64OrDefault]'s result to *p.
func Float64VarOrDefault(s *VarSet, p *float64, key string, defaultValue float64) {
	varFunc(s, p, key, func(key string) (float64, error) { return Float64OrDefault(key, defaultValue) })
}

// DurationVar registers key with s, assigning [Duration]'s result to *p.
func DurationVar(s *VarSet, p *time.Duration, key string) {
	varFunc(s, p, key, Duration)
}

// DurationVarOrDefault registers key with s, assigning [DurationOrDefault]'s
// result to *p.
func DurationVarOrDefault(s *VarSet, p *time.Duration, key string, defaultValue time.Duration) {
	varFunc(s, p, key, func(key string) (time.Duration, error) { return DurationOrDefault(key, defaultValue) })
}

// Parse runs every registered action in registration order. If any fail, it
// returns a [VarSetParseError] aggregating all of them.
func (s *VarSet) Parse() error {
	var errs map[string]error

	for _, a := range s.actions {
		if err := a.run(); err != nil {
			if errs == nil {
				errs = make(map[string]error)
			}

			errs[a.key] = err
		}
	}

	if len(errs) > 0 {
		return VarSetParseError{Errors: errs}
	}

	return nil
}
