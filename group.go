package envvar

// Group provides an easy way to parse multiple environment variables into Go variables or structures.
type Group struct {
	actions map[string]func() error
}

// BoolVar registers a boolean variable with the group. It takes a pointer to
// a bool, a key, and optional configuration functions. The registered action
// retrieves the boolean value associated with the key and assigns it to the
// provided pointer.
func (grp *Group) BoolVar(p *bool, key string, options ...func(c *parseConfig[bool])) {
	grp.actions[key] = func() error {
		value, err := Bool(key, options...)
		if err != nil {
			return err
		}

		*p = value

		return nil
	}
}

// StringVar registers a string variable with the group. It takes a pointer to
// a string, a key, and optional configuration functions. The registered action
// retrieves the string value associated with the key and assigns it to the
// provided pointer.
func (grp *Group) StringVar(p *string, key string, options ...func(c *parseConfig[string])) {
	grp.actions[key] = func() error {
		value, err := String(key, options...)
		if err != nil {
			return err
		}

		*p = value

		return nil
	}
}

// IntVar registers a int variable with the group. It takes a pointer to
// a int, a key, and optional configuration functions. The registered action
// retrieves the int value associated with the key and assigns it to the
// provided pointer.
func (grp *Group) IntVar(p *int, key string, options ...func(c *parseConfig[int])) {
	grp.actions[key] = func() error {
		value, err := Int(key, options...)
		if err != nil {
			return err
		}

		*p = value

		return nil
	}
}

// Parse executes all registered actions in the group. If an error occurs during the parsing
// process, it returns a [GroupParseError] containing all errors.
func (grp *Group) Parse() error {
	errors := make(map[string]error)

	for key, action := range grp.actions {
		err := action()
		if err != nil {
			errors[key] = err
		}
	}

	if len(errors) > 0 {
		return GroupParseError{errors}
	}

	return nil
}

// NewGroup creates a new [Group].
func NewGroup() *Group {
	return &Group{actions: make(map[string]func() error)}
}
