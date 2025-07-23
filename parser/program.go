package parser

// Top level, a list of expressions

type Program struct {
	expressions []Node
}

func (prog *Program) String() string {
	return "Program"
}

// Interface Node
func (prog *Program) Children() []Node {
	return prog.expressions
}

func (prog *Program) SyntaxCheck() error {
	return nil
}

func (prog *Program) Evaluate() (Node, error) {
	return nil, nil
}
