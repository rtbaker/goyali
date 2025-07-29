package lisp

// Top level, a list of expressions

type Program struct {
	Expressions []Node
}

func (prog *Program) String() string {
	return "Program"
}

// Interface Node
func (prog *Program) QuotedValue() Node {
	return NewAtom("", 0, 0)
}

func (prog *Program) Line() int {
	return 0
}

func (prog *Program) Position() int {
	return 0
}

func (prog *Program) Children() []Node {
	return prog.Expressions
}

func (prog *Program) SyntaxCheck() error {
	return nil
}

func (prog *Program) Evaluate() (Node, error) {
	return nil, nil
}
