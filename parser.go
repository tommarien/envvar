package envvar

// Parser provides an easy way to parse environment variables into Go variables or structures.
type Parser struct {
	actions map[string]func() error
}

// BoolVar registers a boolean variable with the parser. It takes a pointer to
// a bool, a key, and optional configuration functions. The registered action
// retrieves the boolean value associated with the key and assigns it to the
// provided pointer.
func (parser *Parser) BoolVar(p *bool, key string, options ...func(c *config[bool])) {
	parser.actions[key] = func() error {
		value, err := Bool(key, options...)
		if err != nil {
			return err
		}

		*p = value

		return nil
	}
}

// StringVar registers a string variable with the parser. It takes a pointer to
// a string, a key, and optional configuration functions. The registered action
// retrieves the string value associated with the key and assigns it to the
// provided pointer.
func (parser *Parser) StringVar(p *string, key string, options ...func(c *config[string])) {
	parser.actions[key] = func() error {
		value, err := String(key, options...)
		if err != nil {
			return err
		}

		*p = value

		return nil
	}
}

// IntVar registers a int variable with the parser. It takes a pointer to
// a int, a key, and optional configuration functions. The registered action
// retrieves the int value associated with the key and assigns it to the
// provided pointer.
func (parser *Parser) IntVar(p *int, key string, options ...func(c *config[int])) {
	parser.actions[key] = func() error {
		value, err := Int(key, options...)
		if err != nil {
			return err
		}

		*p = value

		return nil
	}
}

// Parse executes all registered actions. If an error occurs during the parsing
// process, it returns a [ParserError] containing all errors.
func (parser *Parser) Parse() error {
	errors := make(map[string]error)

	for key, action := range parser.actions {
		err := action()
		if err != nil {
			errors[key] = err
		}
	}

	if len(errors) > 0 {
		return ParserError{errors}
	}

	return nil
}

// NewParser creates a new [Parser].
func NewParser() *Parser {
	return &Parser{actions: make(map[string]func() error)}
}
