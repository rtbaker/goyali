package lisp

// Top level, a list of expressions

type Program struct {
	env         *Env
	Expressions []Node
}

func NewProgram() *Program {
	// top level symbol table
	env := NewEnv(nil)

	// Add the builtin functions here
	env.addSymbol("quote", NewQuoteOp(0, 0))
	env.addSymbol("atom", NewAtomOp(0, 0))

	return &Program{
		env: env,
	}
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

func (prog *Program) AppendNode(n Node) {
	prog.Expressions = append(prog.Expressions, n)
}

func (prog *Program) SyntaxCheck() error {
	return nil
}

func (prog *Program) Evaluate() (Node, error) {
	return nil, nil
}
