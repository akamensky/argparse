package argparse

type Parser interface {
	Parse([]string) error
}

type parser struct{}

func (p *parser) Parse(args []string) error {
	//TODO implement me
	panic("implement me")
}

func New() Parser {
	return &parser{}
}
