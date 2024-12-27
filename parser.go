package envvar

type ParserError struct {
	errors map[string]error
}

func (p ParserError) Error() string {
	panic("unimplemented")
}

func (p ParserError) Unwrap() []error {
	panic("unimplemented")
}

type Parser struct {
	actions map[string]func() error
}

func (parser *Parser) String(p *string, key string) *Parser {
	panic("unimplemented")
}

func (parser *Parser) StringOrDefault(p *string, key, defaultValue string) *Parser {
	panic("unimplemented")
}

func (parser *Parser) Parse() error {
	panic("unimplemented")
}

func NewParser() *Parser {
	return &Parser{actions: make(map[string]func() error)}
}
