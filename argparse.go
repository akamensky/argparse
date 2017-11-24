package argparse

func NewParser(name string, description string) parser {
	p := parser{}

	p.name = name
	p.description = description

	p.args = make(map[string]arg)
	p.commands = make([]command, 0)

	return p
}
